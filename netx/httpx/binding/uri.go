package binding

import (
	"github.com/go-playground/form/v4"
	"net/url"
)

func Uri(m map[string][]string, obj any, tag string) error {
	decoder := form.NewDecoder()
	decoder.SetTagName(tag)
	return decoder.Decode(obj, url.Values(m))
}
