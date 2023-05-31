package binding

import (
	"github.com/go-leo/gox/encodingx/formx"
	"net/http"
	"net/url"
)

func Header(req *http.Request, obj any, tag string) error {
	return formx.Unmarshal(url.Values(req.Header), obj, tag)
}
