package strconvx

import "testing"

// BenchmarkQuote-8        39137730                30.75 ns/op           16 B/op          1 allocs/op
func BenchmarkQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Quote("Hello, World!", "\"")
	}
}

// BenchmarkQuoteV2-8       9224214               127.9 ns/op            80 B/op          5 allocs/op
func BenchmarkQuoteV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteV2("Hello, World!", "\"")
	}
}

// BenchmarkQuoteV3-8      38127211                30.92 ns/op           16 B/op          1 allocs/op
func BenchmarkQuoteV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quoteV3("Hello, World!", "\"")
	}
}
