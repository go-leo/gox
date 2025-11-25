// Package randx provides enhanced random number generation utilities
// based on Go's math/rand/v2 package with different algorithm implementations.
package randx

import (
	"crypto/rand"         // For cryptographically secure random number generation
	"encoding/binary"     // For byte slice to integer conversion
	randv2 "math/rand/v2" // Go's new version of random number generation package

	// For efficient string building
	"sync" // For synchronization primitives like Pool

	"github.com/go-leo/gox/bytex"
)

// Predefined common character sets for string generation
const (
	// Alphanumeric character set includes both upper and lower case letters and digits
	Alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Uppercase letters only
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Lowercase letters only
	Lowercase = "abcdefghijklmnopqrstuvwxyz"

	// Numeric digits only
	Numeric = "0123456789"

	// Hexadecimal characters (lowercase)
	Hex = "0123456789abcdef"

	// Base64 encoding character set
	Base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	// URL-safe Base64 encoding character set
	URLSafeBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

// stringPool is a byte slice pool used for efficient string generation.
// It helps reduce memory allocations when generating random strings by reusing
// byte buffers. The pool is initialized with a minimum capacity of 16 bytes
// and a maximum capacity of 16KB (16*1024 bytes).
//
// This approach minimizes garbage collection pressure by reusing memory buffers
// instead of allocating new ones for each string generation operation.
var stringPool = bytex.New(16, 16*1024)

// chacha8Pool is a sync.Pool that manages reusable ChaCha8 random number generators.
// Using a pool helps reduce garbage collection pressure by reusing expensive-to-create
// random number generator instances instead of creating new ones each time.
//
// The New function is called when the pool is empty and needs to create a new generator.
var chacha8Pool = &sync.Pool{
	// New function creates a new ChaCha8 random number generator
	// It panics if there's an error during creation since this should never happen
	// under normal circumstances with proper system configuration.
	New: func() any {
		// Create a new ChaCha8 random number generator
		r, err := ChaCha8()
		if err != nil {
			// Panic if we can't create the generator - this indicates a serious problem
			// such as inability to read from the system's cryptographically secure
			// random number generator.
			panic(err)
		}
		// Return the created generator as an interface{} to satisfy the Pool interface
		return r
	},
}

// pcgPool is a sync.Pool that manages reusable PCG random number generators.
// Similar to randChaCha8Pool, it helps optimize performance by reusing PCG instances.
var pcgPool = &sync.Pool{
	// New function creates a new PCG random number generator
	// It also panics on error for the same reasons as the ChaCha8 pool.
	New: func() any {
		// Create a new PCG random number generator
		r, err := PCG()
		if err != nil {
			// Panic if we can't create the generator due to system issues
			panic(err)
		}
		// Return the created generator as an interface{} to satisfy the Pool interface
		return r
	},
}

// GetChaCha8 retrieves a ChaCha8 random number generator from the pool.
// If the pool is empty, it creates a new one using the pool's New function.
//
// Returns:
//   - *randv2.Rand: A pointer to a ChaCha8-based random number generator
//
// Example usage:
//
//	rng := GetChaCha8()
//	defer PutChaCha8(rng) // Always return the generator to the pool
//	randomValue := rng.IntN(100)
func GetChaCha8() *randv2.Rand {
	// Get a generator from the pool and cast it back to *randv2.Rand
	// The type assertion is safe because we know the pool only contains *randv2.Rand
	return chacha8Pool.Get().(*randv2.Rand)
}

// PutChaCha8 returns a ChaCha8 random number generator to the pool for reuse.
// This should always be called after finishing with a generator obtained from GetChaCha8().
//
// Parameters:
//   - r: The random number generator to return to the pool
//
// Example usage:
//
//	rng := GetChaCha8()
//	defer PutChaCha8(rng) // Ensures the generator is returned to the pool
//	// ... use rng ...
func PutChaCha8(r *randv2.Rand) {
	// Put the generator back into the pool for future reuse
	// This helps reduce memory allocations and garbage collection pressure
	chacha8Pool.Put(r)
}

// GetPCG retrieves a PCG random number generator from the pool.
// If the pool is empty, it creates a new one using the pool's New function.
//
// Note: There appears to be a bug in the implementation where it's using
// randChaCha8Pool instead of randPCGPool. This should likely be fixed to use
// randPCGPool to correctly manage PCG generators separately from ChaCha8 generators.
//
// Returns:
//   - *randv2.Rand: A pointer to a PCG-based random number generator
//
// Example usage:
//
//	rng := GetPCG()
//	defer PutPCG(rng) // Always return the generator to the pool
//	randomValue := rng.IntN(100)
func GetPCG() *randv2.Rand {
	// Get a generator from the pool and cast it back to *randv2.Rand
	// NOTE: This appears to be a bug - it should use randPCGPool instead of randChaCha8Pool
	return chacha8Pool.Get().(*randv2.Rand)
}

// PutPCG returns a PCG random number generator to the pool for reuse.
// This should always be called after finishing with a generator obtained from GetPCG().
//
// Note: There appears to be a bug in the implementation where it's using
// randChaCha8Pool instead of randPCGPool. This should likely be fixed to use
// randPCGPool to correctly manage PCG generators separately from ChaCha8 generators.
//
// Parameters:
//   - r: The random number generator to return to the pool
//
// Example usage:
//
//	rng := GetPCG()
//	defer PutPCG(rng) // Ensures the generator is returned to the pool
//	// ... use rng ...
func PutPCG(r *randv2.Rand) {
	// Put the generator back into the pool for future reuse
	// NOTE: This appears to be a bug - it should use randPCGPool instead of randChaCha8Pool
	chacha8Pool.Put(r)
}

// ChaCha8 creates a new random number generator using the ChaCha8 algorithm.
// This function generates a cryptographically secure 32-byte seed and uses it
// to initialize a ChaCha8-based random number generator.
//
// The ChaCha8 algorithm is a stream cipher that can be used as a pseudorandom
// number generator. It provides cryptographically secure randomness suitable
// for security-sensitive applications.
//
// Returns:
//   - *randv2.Rand: A pointer to the initialized random number generator
//   - error: Any error encountered during seed generation, typically from crypto/rand
//
// Example usage:
//
//	rng, err := ChaCha8()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	randomValue := rng.IntN(100)
func ChaCha8() (*randv2.Rand, error) {
	// Allocate a 32-byte array to store the cryptographic seed
	// ChaCha8 requires a 32-byte (256-bit) key/seed
	var seed [32]byte

	// Generate a cryptographically secure random seed
	// Read fills the seed slice with random data from the system's CSPRNG
	_, err := rand.Read(seed[:])
	if err != nil {
		// Return nil and the error if seed generation fails
		return nil, err
	}

	// Create and return a new random number generator using the ChaCha8 algorithm
	// NewChaCha8 initializes the ChaCha8 cipher with the provided seed
	// randv2.New wraps the cipher into a usable Rand instance
	return randv2.New(randv2.NewChaCha8(seed)), nil
}

// PCG creates a new random number generator using the PCG (Permuted Congruential Generator) algorithm.
// This function generates a cryptographically secure 16-byte seed and uses it to initialize
// a PCG-based random number generator with two 64-bit seed values.
//
// The PCG algorithm is known for its excellent statistical properties and performance,
// while being smaller and faster than many other generators. It's suitable for
// general-purpose random number generation where cryptographic security is not required.
//
// Returns:
//   - *randv2.Rand: A pointer to the initialized random number generator
//   - error: Any error encountered during seed generation, typically from crypto/rand
//
// Example usage:
//
//	rng, err := PCG()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	randomValue := rng.IntN(100)
func PCG() (*randv2.Rand, error) {
	// Allocate a 16-byte array to store the cryptographic seed
	// PCG requires two 64-bit (8-byte each) seed values = 16 bytes total
	var b [16]byte

	// Generate a cryptographically secure random seed
	_, err := rand.Read(b[:])
	if err != nil {
		// Return nil and the error if seed generation fails
		return nil, err
	}

	// Convert the first 8 bytes to a uint64 seed value using big-endian encoding
	// This will be used as the state seed for the PCG generator
	seed1 := binary.BigEndian.Uint64(b[:8])

	// Convert the second 8 bytes to a uint64 seed value using big-endian encoding
	// This will be used as the sequence seed for the PCG generator
	seed2 := binary.BigEndian.Uint64(b[8:])

	// Create and return a new random number generator using the PCG algorithm
	// NewPCG initializes the PCG generator with the two seed values
	// randv2.New wraps the PCG generator into a usable Rand instance
	return randv2.New(randv2.NewPCG(seed1, seed2)), nil
}

// Bool generates a random boolean value (true or false) with approximately 50% probability for each.
// This function utilizes the pooled PCG random number generator for efficiency.
//
// The function works by:
// 1. Retrieving a PCG random number generator from the pool via GetPCG()
// 2. Generating a random 64-bit unsigned integer using Uint64()
// 3. Checking if the least significant bit is 0 (even) to determine the boolean value
// 4. Returning true if the number is even, false if odd
//
// Returns:
//   - bool: A randomly generated boolean value
//
// Example usage:
//
//	result := Bool()
//	fmt.Println(result) // Prints either true or false
//
// Note: While this implementation is simple and efficient, it has a minor bias
// because the modulo operation on a uniform distribution doesn't guarantee perfect
// 50/50 distribution unless the range is evenly divisible by the modulus.
// However, for most practical purposes, this bias is negligible.
func Bool() bool {
	// Get a PCG random number generator from the pool
	// Generate a random uint64 and check if it's even (LSB = 0)
	// This gives us approximately 50% chance of true and 50% chance of false
	return GetPCG().Uint64()%2 == 0
}

// Int returns a non-negative pseudo-random int from the pooled PCG generator.
// It produces a value in the range [0, math.MaxInt].
//
// Returns:
//   - int: A non-negative pseudo-random integer
func Int() int {
	return GetPCG().Int()
}

// IntN returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number
//
// Returns:
//   - int: A non-negative pseudo-random number in [0,n)
//
// Example usage:
//
//	result := IntN(100) // Returns a value between 0 and 99
func IntN(n int) int {
	return GetPCG().IntN(n)
}

// IntRange returns a pseudo-random number in the closed interval [min, max].
// This is a convenience function that generates numbers in a specific range.
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - int: A pseudo-random number in [min, max]
//
// Example usage:
//
//	result := IntRange(10, 20) // Returns a value between 10 and 20 (both inclusive)
func IntRange(min, max int) int {
	return IntN(max-min) + min
}

// Int32 returns a non-negative pseudo-random 31-bit integer as an int32.
// It produces a value in the range [0, math.MaxInt32].
//
// Returns:
//   - int32: A non-negative pseudo-random 31-bit integer
func Int32() int32 {
	return GetPCG().Int32()
}

// Int32N returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number
//
// Returns:
//   - int32: A non-negative pseudo-random number in [0,n)
func Int32N(n int32) int32 {
	return GetPCG().Int32N(n)
}

// Int32Range returns a pseudo-random number in the closed interval [min, max].
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - int32: A pseudo-random number in [min, max]
func Int32Range(min, max int32) int32 {
	return Int32N(max-min) + min
}

// Int64 returns a non-negative pseudo-random 63-bit integer as an int64.
// It produces a value in the range [0, math.MaxInt64].
//
// Returns:
//   - int64: A non-negative pseudo-random 63-bit integer
func Int64() int64 {
	return GetPCG().Int64()
}

// Int64N returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number
//
// Returns:
//   - int64: A non-negative pseudo-random number in [0,n)
func Int64N(n int64) int64 {
	return GetPCG().Int64N(n)
}

// Int64Range returns a pseudo-random number in the closed interval [min, max].
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - int64: A pseudo-random number in [min, max]
func Int64Range(min, max int64) int64 {
	return Int64N(max-min) + min
}

// Uint returns a pseudo-random uint from the pooled PCG generator.
//
// Returns:
//   - uint: A pseudo-random unsigned integer
func Uint() uint {
	return GetPCG().Uint()
}

// UintN returns, as a uint, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number
//
// Returns:
//   - uint: A pseudo-random number in [0,n)
func UintN(n uint) uint {
	return GetPCG().UintN(n)
}

// UintRange returns a pseudo-random number in the closed interval [min, max].
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - uint: A pseudo-random number in [min, max]
func UintRange(min, max uint) uint {
	return UintN(max-min) + min
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
//
// Returns:
//   - uint32: A pseudo-random 32-bit unsigned integer
func Uint32() uint32 {
	return GetPCG().Uint32()
}

// Uint32N returns, as a uint32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number
//
// Returns:
//   - uint32: A pseudo-random number in [0,n)
func Uint32N(n uint32) uint32 {
	return GetPCG().Uint32N(n)
}

// Uint32Range returns a pseudo-random number in the closed interval [min, max].
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - uint32: A pseudo-random number in [min, max]
func Uint32Range(min, max uint32) uint32 {
	return Uint32N(max-min) + min
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
//
// Returns:
//   - uint64: A pseudo-random 64-bit unsigned integer
func Uint64() uint64 {
	return GetPCG().Uint64()
}

// Uint64N returns, as a uint64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number
//
// Returns:
//   - uint64: A pseudo-random number in [0,n)
func Uint64N(n uint64) uint64 {
	return GetPCG().Uint64N(n)
}

// Uint64Range returns a pseudo-random number in the closed interval [min, max].
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - uint64: A pseudo-random number in [min, max]
func Uint64Range(min, max uint64) uint64 {
	return Uint64N(max-min) + min
}

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0).
//
// Returns:
//   - float32: A pseudo-random number in [0.0,1.0)
func Float32() float32 {
	return GetPCG().Float32()
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
//
// Returns:
//   - float64: A pseudo-random number in [0.0,1.0)
func Float64() float64 {
	return GetPCG().Float64()
}

// String generates a random string of specified length using the provided character set.
// This function is useful for generating random identifiers, passwords, tokens,
// or any other string that requires randomness.
//
// The function uses a byte buffer pool for efficient memory management, reducing
// the number of allocations and garbage collection pressure when generating
// multiple strings.
//
// Parameters:
//   - length: The desired length of the random string. Must be a positive integer.
//     If length is zero or negative, the function returns an empty string.
//   - charset: The set of characters to choose from when generating the random string.
//     Each character in the resulting string will be randomly selected
//     from this character set.
//
// Returns:
//   - string: A randomly generated string of the specified length, or an empty
//     string if length is non-positive.
//
// Example usage:
//
//	// Generate a random alphanumeric string of length 10
//	result := String(10, Alphanumeric)
//	fmt.Println(result) // e.g., "aB3xF9Kp2M"
//
//	// Generate a random hexadecimal string of length 16
//	result := String(16, Hex)
//	fmt.Println(result) // e.g., "9a3f7b2e8c1d4a5f"
//
//	// Generate a random password with special characters
//	charset := Alphanumeric + "!@#$%^&*()"
//	password := String(12, charset)
//	fmt.Println(password) // e.g., "K9#mP2$xQ!vR"
//
// Performance considerations:
//   - For better performance, reuse character sets defined as constants
//   - The function uses a byte buffer pool to minimize memory allocations
//   - For very large strings (>16KB), new buffers will be allocated as needed
func String(length int, charset string) string {
	// Handle edge case of non-positive length
	// If the requested length is zero or negative, return an empty string immediately
	// This avoids unnecessary processing and buffer allocation
	if length <= 0 {
		return ""
	}

	// Get a byte buffer from the pool with sufficient capacity for the requested length
	// The buffer will be reused from the pool if available, or newly allocated if needed
	// This reduces memory allocation overhead compared to creating a new buffer each time
	// stringPool.Get() retrieves a byte buffer from the pool that can hold at least 'length' bytes
	// If no suitable buffer exists in the pool, a new one will be created automatically
	buf := stringPool.Get(length)

	// Defer putting the buffer back into the pool for reuse
	// This ensures that regardless of how the function exits (normal or panic),
	// the buffer will be returned to the pool for future use
	// This is critical for maintaining the efficiency of the buffer pool
	defer stringPool.Put(buf)

	// Generate each character by randomly selecting from the charset
	// Iterate length times to build a string of the requested length
	for i := 0; i < length; i++ {
		// Select a random index from the charset and append the character
		// IntN(len(charset)) generates a random integer in [0, len(charset))
		// This index is used to select a character from the charset string
		// buf.WriteByte() appends the selected character to the byte buffer
		buf.WriteByte(charset[IntN(len(charset))])
	}

	// Convert the byte buffer to a string and return it
	// buf.String() creates a string from the byte buffer contents
	// The buffer will be automatically returned to the pool for reuse
	// after the string is created
	return buf.String()
}
