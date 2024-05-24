package convx

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i any) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i any) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToFloat64E casts an interface to a float64 type.
func ToFloat64E(i any) (float64, error) {
	return ToFloatE[float64](i)
}

// ToFloat32E casts an interface to a float32 type.
func ToFloat32E(i any) (float32, error) {
	return ToFloatE[float32](i)
}
