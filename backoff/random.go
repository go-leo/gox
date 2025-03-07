package backoff

import (
	"context"
	"math/rand"
	"time"
)

// Random generates a function that waits for a random time in the range [alpha, beta) between calls.
func Random(alpha, beta time.Duration) Func {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(ctx context.Context, attempt uint) time.Duration {
		return alpha + time.Duration(r.Int63n(beta.Nanoseconds()-alpha.Nanoseconds()))
	}
}
