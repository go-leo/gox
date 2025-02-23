package binding

import (
	"github.com/go-playground/form/v4"
	"net/http"
)

func Query(req *http.Request, obj any, tag string) error {
	decoder := form.NewDecoder()
	decoder.SetTagName(tag)
	return decoder.Decode(obj, req.URL.Query())
}
