package binding

import (
	"github.com/go-leo/gox/encodingx/tomlx"
	"net/http"
)

func TOML(req *http.Request, obj any) error {
	return tomlx.NewDecoder(req.Body).Decode(obj)
}
