package binding

import (
	"github.com/go-playground/form/v4"
	"net/http"
	"net/url"
)

func Header(req *http.Request, obj any, tag string) error {
	decoder := form.NewDecoder()
	decoder.SetTagName(tag)
	return decoder.Decode(obj, url.Values(req.Header))
}
