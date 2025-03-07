package backoff

import (
	"context"
	"crypto/rand"
	randv2 "math/rand/v2"
	"time"
)

var jitterRand *randv2.Rand

func init() {
	seed := [32]byte{}
	_, err := rand.Read(seed[:])
	if err != nil {
		panic(err)
	}
	jitterRand = randv2.New(randv2.NewChaCha8(seed))
}

// JitterUp adds random jitter to the interval.
//
// This adds or subtracts time from the interval within a given jitter fraction.
// For example for 10s and jitter 0.1, it will return a time within [9s, 11s])
func JitterUp(backoff Func, jitter float64) Func {
	return func(ctx context.Context, attempt uint) time.Duration {
		interval := backoff(ctx, attempt)
		return jitterUp(interval, jitter)
	}
}

func jitterUp(interval time.Duration, jitter float64) time.Duration {
	multiplier := jitter * (jitterRand.Float64()*2 - 1)
	return time.Duration(float64(interval) * (1 + multiplier))
}

func JitterUpFactory(factory Factory, jitter float64) Factory {
	return func(delta time.Duration) Func {
		return JitterUp(factory(delta), jitter)
	}
}
