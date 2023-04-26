//go:build !go1.20

package contextx

import (
	"context"
)

func Error(ctx context.Context) error {
	return ctx.Err()
}
