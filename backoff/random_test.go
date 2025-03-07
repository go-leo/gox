package backoff

import (
	"context"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	random := Random(time.Second, time.Minute)
	duration := random(context.Background(), 0)
	t.Log(duration)
}
