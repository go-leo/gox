package queryx

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

func Marshal(v any) (url.Values, error) {
	return query.Values(v)
}
