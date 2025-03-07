package backoff

import (
	"context"
	"testing"
	"time"
)

func TestJitterUp(t *testing.T) {
	random := JitterUp(Constant(time.Second), 0.1)
	duration := random(context.Background(), 0)
	t.Log(duration)
}
