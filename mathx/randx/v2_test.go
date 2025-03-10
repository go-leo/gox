package randx

import "testing"

func TestNewChacha8(t *testing.T) {
	rand, err := NewChaCha8()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		t.Log(rand.Uint64())
	}
}

func TestNewPCG(t *testing.T) {
	rand, err := NewPCG()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		t.Log(rand.Uint64())
	}
}
