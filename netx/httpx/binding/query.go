package binding

import (
	"github.com/go-leo/gox/encodingx/formx"
	"net/http"
)

func Query(req *http.Request, obj any, tag string) error {
	return formx.Unmarshal(req.URL.Query(), obj, tag)
}
