package binding

import (
	"github.com/go-leo/gox/encodingx/yamlx"
	"net/http"
)

func YAML(req *http.Request, obj any) error {
	return yamlx.NewDecoder(req.Body).Decode(obj)
}
