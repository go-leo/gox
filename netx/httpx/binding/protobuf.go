package binding

import (
	"github.com/go-leo/gox/encodingx/protobufx"
	"net/http"
)

func ProtoBuf(req *http.Request, obj any) error {
	return protobufx.NewDecoder(req.Body).Decode(obj)
}
