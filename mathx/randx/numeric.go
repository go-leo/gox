package randx

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func Int63() int64 {
	r := Get()
	i := r.Int63()
	Put(r)
	return i
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func Uint32() uint32 {
	r := Get()
	i := r.Uint32()
	Put(r)
	return i
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func Uint64() uint64 {
	r := Get()
	i := r.Uint64()
	Put(r)
	return i
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func Int31() int32 {
	r := Get()
	i := r.Int31()
	Put(r)
	return i
}

// Int returns a non-negative pseudo-random int.
func Int() int {
	r := Get()
	i := r.Int()
	Put(r)
	return i
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Int63n(n int64) int64 {
	r := Get()
	i := r.Int63n(n)
	Put(r)
	return i
}

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Int31n(n int32) int32 {
	r := Get()
	i := r.Int31n(n)
	Put(r)
	return i
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	r := Get()
	i := r.Intn(n)
	Put(r)
	return i
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func Float64() float64 {
	r := Get()
	f := r.Float64()
	Put(r)
	return f
}

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0).
func Float32() float32 {
	r := Get()
	f := r.Float32()
	Put(r)
	return f
}

func PickItemsInt32(size int, n ...int32) []int32 {
	if size <= 0 {
		return []int32{}
	}
	copy := append([]int32{}, n...)
	if size >= len(n) {
		return copy
	}
	indexSet := make(map[int]struct{})
	for i := 0; i < size; i++ {
		index := Intn(len(n))
		if _, ok := indexSet[index]; !ok {
			indexSet[index] = struct{}{}
			continue
		}
		for {
			index = Intn(len(n))
			if _, ok := indexSet[index]; !ok {
				indexSet[index] = struct{}{}
				break
			}
		}

	}

	return copy[:size]
}
