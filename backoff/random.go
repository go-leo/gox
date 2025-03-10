package backoff

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	randv2 "math/rand/v2"
	"time"
)

var randomRand *randv2.Rand

func init() {
	var err error
	randomRand, err = randx.NewChaCha8()
	if err != nil {
		panic(err)
	}
}

// Random generates a function that waits for a random time in the range [alpha, beta) between calls.
func Random(alpha, beta time.Duration) Func {
	return func(ctx context.Context, attempt uint) time.Duration {
		return alpha + time.Duration(randomRand.Int64N(beta.Nanoseconds()-alpha.Nanoseconds()))
	}
}

func RandomFactory(alpha, beta time.Duration) Factory {
	return func(delta time.Duration) Func {
		return func(ctx context.Context, attempt uint) time.Duration {
			return Random(alpha, beta)(ctx, attempt)
		}
	}
}
