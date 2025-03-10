package randx

import (
	"crypto/rand"
	"encoding/binary"
	randv2 "math/rand/v2"
)

func NewChaCha8() (*randv2.Rand, error) {
	var seed [32]byte
	_, err := rand.Read(seed[:])
	if err != nil {
		return nil, err
	}
	return randv2.New(randv2.NewChaCha8(seed)), nil
}

func NewPCG() (*randv2.Rand, error) {
	var b [16]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return nil, err
	}
	seed1 := binary.BigEndian.Uint64(b[:8])
	seed2 := binary.BigEndian.Uint64(b[8:])
	return randv2.New(randv2.NewPCG(seed1, seed2)), nil
}
