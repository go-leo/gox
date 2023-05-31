package binding

import (
	"github.com/go-leo/gox/encodingx/xmlx"
	"net/http"
)

func XML(req *http.Request, obj any) error {
	return xmlx.NewDecoder(req.Body).Decode(obj)
}
