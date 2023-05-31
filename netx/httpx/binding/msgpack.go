package binding

import (
	"github.com/go-leo/gox/encodingx/msgpackx"
	"net/http"
)

func MsgPack(req *http.Request, obj any) error {
	return msgpackx.NewDecoder(req.Body).Decode(obj)
}
