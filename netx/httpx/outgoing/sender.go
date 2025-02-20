package outgoing

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/encodingx/jsonx"
	"github.com/go-leo/gox/encodingx/queryx"
	"github.com/go-leo/gox/encodingx/xmlx"
	"github.com/go-leo/gox/iox"
	"github.com/go-leo/gox/netx/httpx"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/gox/stringx"
	"google.golang.org/protobuf/proto"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

var (
	ErrMethodEmpty = errors.New("method is empty")
	ErrURLEmpty    = errors.New("url is empty")
)

type MarshalError struct {
	Object any
	Err    error
}

func (e MarshalError) Error() string {
	return fmt.Sprintf("failed to marshal body")
}

type RequestSender interface {
	Method(method string) URLSender
	Get() URLSender
	Head() URLSender
	Post() URLSender
	Put() URLSender
	Patch() URLSender
	Delete() URLSender
	Connect() URLSender
	Options() URLSender
	Trace() URLSender
}

type URLSender interface {
	URL(uri *url.URL) PayloadSender
	URLString(urlString string) PayloadSender
}

type FormData struct {
	FieldName string
	Value     string
	File      io.Reader
	Filename  string
}

type PayloadSender interface {
	// Query methods
	Query(name, value string) PayloadSender
	AddQuery(key, value string) PayloadSender
	DelQuery(name string) PayloadSender
	QueryString(q string) PayloadSender
	Queries(queries url.Values) PayloadSender
	QueryObject(q any) PayloadSender

	// Header methods
	Header(name, value string, uncanonical ...bool) PayloadSender
	AddHeader(name, value string, uncanonical ...bool) PayloadSender
	DelHeader(name string) PayloadSender
	Headers(header http.Header) PayloadSender
	UserAgent(ua string) PayloadSender

	// Auth methods
	BasicAuth(username, password string) PayloadSender
	BearerAuth(token string) PayloadSender
	CustomAuth(scheme, token string) PayloadSender

	// Cache-Control methods
	CacheControl(directives ...string) PayloadSender
	IfModifiedSince(t time.Time) PayloadSender
	IfUnmodifiedSince(t time.Time) PayloadSender
	IfNoneMatch(etag string) PayloadSender
	IfMatch(etags ...string) PayloadSender

	// Cookie methods
	Cookie(cookie *http.Cookie) PayloadSender
	AddCookie(cookie *http.Cookie) PayloadSender
	DelCookie(cookie *http.Cookie) PayloadSender
	Cookies(cookies ...*http.Cookie) PayloadSender

	// Body methods
	Body(body io.Reader, contentType string) PayloadSender
	BytesBody(body []byte, contentType string) PayloadSender
	TextBody(body string, contentType string) PayloadSender
	FormBody(form url.Values) PayloadSender
	FormObjectBody(body any) PayloadSender
	ObjectBody(body any, marshal func(any) ([]byte, error), contentType string) PayloadSender
	JSONBody(body any) PayloadSender
	XMLBody(body any) PayloadSender
	ProtobufBody(body proto.Message) PayloadSender
	GobBody(body any) PayloadSender
	MultipartBody(formData ...*FormData) PayloadSender

	// other methods
	Middleware(middlewares ...Middleware) PayloadSender
	Build(ctx context.Context) (*http.Request, error)
	Send(ctx context.Context, clis ...*http.Client) (ResponseReceiver, error)
}

type sender struct {
	err         error
	method      string
	uri         *url.URL
	queries     url.Values
	headers     http.Header
	body        io.Reader
	cookies     map[string][]*http.Cookie
	middlewares []Middleware
}

func (s *sender) Method(method string) URLSender {
	if s.err != nil {
		return s
	}
	s.method = method
	return s
}

func (s *sender) Get() URLSender {
	return s.Method(http.MethodGet)
}

func (s *sender) Head() URLSender {
	return s.Method(http.MethodHead)
}

func (s *sender) Post() URLSender {
	return s.Method(http.MethodPost)
}

func (s *sender) Put() URLSender {
	return s.Method(http.MethodPut)
}

func (s *sender) Patch() URLSender {
	return s.Method(http.MethodPatch)
}

func (s *sender) Delete() URLSender {
	return s.Method(http.MethodDelete)
}

func (s *sender) Connect() URLSender {
	return s.Method(http.MethodConnect)
}

func (s *sender) Options() URLSender {
	return s.Method(http.MethodOptions)
}

func (s *sender) Trace() URLSender {
	return s.Method(http.MethodTrace)
}

func (s *sender) URL(uri *url.URL) PayloadSender {
	if s.err != nil {
		return s
	}
	s.uri = uri
	return s
}

func (s *sender) URLString(urlString string) PayloadSender {
	if s.err != nil {
		return s
	}
	uri, err := url.Parse(urlString)
	if err != nil {
		s.err = err
		return s
	}
	s.uri = uri
	return s
}

func (s *sender) Query(name, value string) PayloadSender {
	if s.err != nil {
		return s
	}
	s.query().Set(name, value)
	return s
}

func (s *sender) AddQuery(key, value string) PayloadSender {
	if s.err != nil {
		return s
	}
	s.query().Add(key, value)
	return s
}

func (s *sender) DelQuery(name string) PayloadSender {
	if s.err != nil {
		return s
	}
	s.query().Del(name)
	return s
}

func (s *sender) QueryString(q string) PayloadSender {
	if s.err != nil {
		return s
	}
	queries, err := url.ParseQuery(q)
	if err != nil {
		s.err = err
		return s
	}
	return s.Queries(queries)
}

func (s *sender) Queries(queries url.Values) PayloadSender {
	if s.err != nil {
		return s
	}
	for key, values := range queries {
		for _, value := range values {
			s.query().Add(key, value)
		}
	}
	return s
}

func (s *sender) QueryObject(q any) PayloadSender {
	if s.err != nil {
		return s
	}
	values, err := queryx.Marshal(q)
	if err != nil {
		s.err = err
		return s
	}
	return s.Queries(values)
}

func (s *sender) Header(name, value string, uncanonical ...bool) PayloadSender {
	if s.err != nil {
		return s
	}
	if slicex.IsNotEmpty(uncanonical) && uncanonical[0] {
		s.header()[name] = []string{value}
		return s
	}
	s.header().Set(name, value)
	return s
}

func (s *sender) AddHeader(name, value string, uncanonical ...bool) PayloadSender {
	if s.err != nil {
		return s
	}
	if slicex.IsNotEmpty(uncanonical) && uncanonical[0] {
		s.header()[name] = append(s.header()[name], value)
		return s
	}
	s.header().Add(name, value)
	return s
}

func (s *sender) DelHeader(name string) PayloadSender {
	if s.err != nil {
		return s
	}
	s.header().Del(name)
	return s
}

func (s *sender) Headers(header http.Header) PayloadSender {
	if s.err != nil {
		return s
	}
	for key, values := range header {
		for _, value := range values {
			s.header().Add(key, value)
		}
	}
	return s
}

func (s *sender) BasicAuth(username, password string) PayloadSender {
	if s.err != nil {
		return s
	}
	token := base64.StdEncoding.EncodeToString(convx.StringToBytes(username + ":" + password))
	return s.CustomAuth("Basic", token)
}

func (s *sender) BearerAuth(token string) PayloadSender {
	return s.CustomAuth("Bearer", token)
}

func (s *sender) CustomAuth(scheme, token string) PayloadSender {
	return s.Header("Authorization", scheme+" "+token)
}

func (s *sender) UserAgent(ua string) PayloadSender {
	return s.Header("User-Agent", ua)
}

func (s *sender) IfModifiedSince(t time.Time) PayloadSender {
	return s.Header("If-Modified-Since", t.UTC().Format(http.TimeFormat))
}

func (s *sender) IfUnmodifiedSince(t time.Time) PayloadSender {
	return s.Header("If-Unmodified-Since", t.UTC().Format(http.TimeFormat))
}

func (s *sender) IfNoneMatch(etag string) PayloadSender {
	return s.Header("If-None-Match", etag)
}

func (s *sender) IfMatch(etags ...string) PayloadSender {
	return s.Header("If-Match", strings.Join(etags, ", "))
}

func (s *sender) CacheControl(directives ...string) PayloadSender {
	return s.Header("Cache-Control", strings.Join(directives, ", "))
}

func (s *sender) Cookie(cookie *http.Cookie) PayloadSender {
	if s.err != nil {
		return s
	}
	s.cookie()[cookie.Name] = []*http.Cookie{cookie}
	return s
}

func (s *sender) AddCookie(cookie *http.Cookie) PayloadSender {
	return s.Cookies(cookie)
}

func (s *sender) DelCookie(cookie *http.Cookie) PayloadSender {
	if s.err != nil {
		return s
	}
	delete(s.cookie(), cookie.Name)
	return s
}

func (s *sender) Cookies(cookies ...*http.Cookie) PayloadSender {
	if s.err != nil {
		return s
	}
	for _, cookie := range cookies {
		s.cookie()[cookie.Name] = append(s.cookie()[cookie.Name], cookie)
	}
	return s
}

func (s *sender) Body(body io.Reader, contentType string) PayloadSender {
	if s.err != nil {
		return s
	}
	s.body = body
	s.Header("Content-Type", contentType)
	l, ok := iox.Len(body)
	if ok {
		s.Header("Content-Length", convx.ToString(l))
	}
	return s
}

func (s *sender) BytesBody(body []byte, contentType string) PayloadSender {
	return s.Body(bytes.NewReader(body), contentType)
}

func (s *sender) TextBody(body string, contentType string) PayloadSender {
	return s.Body(strings.NewReader(body), contentType)
}

func (s *sender) FormBody(form url.Values) PayloadSender {
	return s.TextBody(form.Encode(), "application/x-www-form-urlencoded")
}

func (s *sender) FormObjectBody(body any) PayloadSender {
	if s.err != nil {
		return s
	}
	form, err := queryx.Marshal(body)
	if err != nil {
		s.err = err
		return s
	}
	return s.FormBody(form)
}

func (s *sender) ObjectBody(body any, marshal func(any) ([]byte, error), contentType string) PayloadSender {
	if s.err != nil {
		return s
	}
	data, err := marshal(body)
	if err != nil {
		s.err = MarshalError{Object: body, Err: err}
		return s
	}
	return s.BytesBody(data, contentType)
}

func (s *sender) JSONBody(body any) PayloadSender {
	marshal := func(v any) ([]byte, error) {
		buffer := &bytes.Buffer{}
		encoder := jsonx.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(v)
		return buffer.Bytes(), err
	}
	return s.ObjectBody(body, marshal, "application/json")
}

func (s *sender) XMLBody(body any) PayloadSender {
	return s.ObjectBody(body, xmlx.Marshal, "application/xml")
}

func (s *sender) ProtobufBody(body proto.Message) PayloadSender {
	marshal := func(v any) ([]byte, error) {
		message, _ := v.(proto.Message)
		return proto.Marshal(message)
	}
	return s.ObjectBody(body, marshal, "application/x-protobuf")
}

func (s *sender) GobBody(body any) PayloadSender {
	marshal := func(v any) ([]byte, error) {
		buffer := &bytes.Buffer{}
		encoder := gob.NewEncoder(buffer)
		if err := encoder.Encode(v); err != nil {
			return nil, err
		}
		return buffer.Bytes(), nil
	}
	return s.ObjectBody(body, marshal, "application/x-gob")
}

func (s *sender) MultipartBody(formData ...*FormData) PayloadSender {
	if s.err != nil {
		return s
	}
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	for _, form := range formData {
		if form.File == nil {
			_ = writer.WriteField(form.FieldName, form.Value)
			continue
		}
		mf, err := writer.CreateFormFile(form.FieldName, filepath.Base(form.Filename))
		if err != nil {
			s.err = err
			return s
		}
		if _, err = io.Copy(mf, form.File); err != nil {
			s.err = err
			return s
		}
	}
	if err := writer.Close(); err != nil {
		s.err = err
		return s
	}
	return s.Body(payload, writer.FormDataContentType())
}

func (s *sender) Middleware(middlewares ...Middleware) PayloadSender {
	s.middlewares = append(s.middlewares, middlewares...)
	return s
}

func (s *sender) Build(ctx context.Context) (*http.Request, error) {
	if s.err != nil {
		return nil, s.err
	}
	if stringx.IsBlank(s.method) {
		return nil, ErrMethodEmpty
	}
	if s.uri == nil {
		return nil, ErrURLEmpty
	}
	query := s.uri.Query()
	for name, values := range s.query() {
		for _, value := range values {
			query.Add(name, value)
		}
	}
	s.uri.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, s.method, s.uri.String(), s.body)
	if err != nil {
		return nil, err
	}
	for key, values := range s.header() {
		req.Header[key] = append(req.Header[key], values...)
	}
	for _, cookies := range s.cookie() {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
	}
	return req, nil
}

func (s *sender) Send(ctx context.Context, clients ...*http.Client) (ResponseReceiver, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var cli *http.Client

	if req, err = s.Build(ctx); err != nil {
		return nil, err
	}
	if len(clients) > 0 && clients[0] != nil {
		cli = clients[0]
	} else {
		cli = httpx.PooledClient()
	}
	middleware := chainMiddlewares(s.middlewares...)
	if middleware == nil {
		if resp, err = invoke(ctx, req, cli); err != nil {
			return nil, err
		}
		return Receiver(resp), nil
	}
	if resp, err = middleware(ctx, req, cli, invoke); err != nil {
		return nil, err
	}
	return Receiver(resp), nil
}

func (s *sender) query() url.Values {
	if s.queries == nil {
		s.queries = make(url.Values)
	}
	return s.queries
}

func (s *sender) header() http.Header {
	if s.headers == nil {
		s.headers = make(http.Header)
	}
	return s.headers
}

func (s *sender) cookie() map[string][]*http.Cookie {
	if s.cookies == nil {
		s.cookies = make(map[string][]*http.Cookie)
	}
	return s.cookies
}

// Sender returns a new RequestSender.
func Sender() RequestSender {
	return new(sender)
}

// Method returns a new RequestSender with the specified method.
func Method(method string) URLSender {
	s := new(sender)
	if s.err != nil {
		return s
	}
	s.method = method
	return s
}

// Get returns a new RequestSender with the method GET.
func Get() URLSender {
	return Sender().Method(http.MethodGet)
}

// Head returns a new RequestSender with the method HEAD.
func Head() URLSender {
	return Sender().Method(http.MethodHead)
}

// Post returns a new RequestSender with the method POST.
func Post() URLSender {
	return Sender().Method(http.MethodPost)
}

// Put returns a new RequestSender with the method PUT.
func Put() URLSender {
	return Sender().Method(http.MethodPut)
}

// Patch returns a new RequestSender with the method PATCH.
func Patch() URLSender {
	return Sender().Method(http.MethodPatch)
}

// Delete returns a new RequestSender with the method DELETE.
func Delete() URLSender {
	return Sender().Method(http.MethodDelete)
}

// Connect returns a new RequestSender with the method CONNECT.
func Connect() URLSender {
	return Sender().Method(http.MethodConnect)
}

// Options returns a new RequestSender with the method OPTIONS.
func Options() URLSender {
	return Sender().Method(http.MethodOptions)
}

// Trace returns a new RequestSender with the method TRACE.
func Trace() URLSender {
	return Sender().Method(http.MethodTrace)
}
