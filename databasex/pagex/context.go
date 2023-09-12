package pagex

import "context"

type key struct{}

func NewContext(ctx context.Context, pageNum uint64, pageSize uint64, opts ...Option) (context.Context, error) {
	page, err := NewPage(pageNum, pageSize, opts...)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, key{}, page), nil
}

func FromContext(ctx context.Context) (*Page, bool) {
	v, ok := ctx.Value(key{}).(*Page)
	return v, ok
}
