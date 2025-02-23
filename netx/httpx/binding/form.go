package binding

import (
	"errors"
	"github.com/go-leo/gox/netx/httpx/binding/internal/multipart"
	"github.com/go-playground/form/v4"
	"net/http"
)

const defaultMemory = 32 << 20

func Form(req *http.Request, obj any, tag string) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	decoder := form.NewDecoder()
	decoder.SetTagName(tag)
	return decoder.Decode(obj, req.Form)
}

func PostForm(req *http.Request, obj any, tag string) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	decoder := form.NewDecoder()
	decoder.SetTagName(tag)
	return decoder.Decode(obj, req.PostForm)
}

func MultipartForm(req *http.Request, obj any, tag string) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	decoder := form.NewDecoder()
	decoder.SetTagName(tag)
	err := decoder.Decode(obj, req.Form)
	if err != nil {
		return err
	}
	return multipart.MappingByPtr(req, obj, tag)
}
