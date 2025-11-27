// File: rand_test.go

package randx

import (
	"strings"
	"testing"
)

// TestStringGeneration 测试随机字符串生成功能
func TestStringGeneration(t *testing.T) {
	length := 10
	charset := Alphanumeric
	
	// 生成随机字符串
	result := String(length, charset)
	
	// 验证长度是否正确
	if len(result) != length {
		t.Errorf("Expected length %d, got %d", length, len(result))
	}
	
	// 验证字符是否都在指定字符集中
	for _, char := range result {
		if !strings.ContainsRune(charset, char) {
			t.Errorf("Character %c not in charset %s", char, charset)
		}
	}
}

// TestBoolGeneration 测试布尔值生成功能
func TestBoolGeneration(t *testing.T) {
	// 多次调用验证返回的是布尔值
	for i := 0; i < 100; i++ {
		result := Bool()
		if result != true && result != false {
			t.Errorf("Expected boolean value, got %v", result)
		}
	}
}

// TestIntGeneration 测试整数生成功能
func TestIntGeneration(t *testing.T) {
	// 调用Int()并验证返回值是非负整数
	result := Int()
	if result < 0 {
		t.Errorf("Expected non-negative integer, got %d", result)
	}
}

// TestIntNGeneration 测试范围内的整数生成功能
func TestIntNGeneration(t *testing.T) {
	n := 100
	for i := 0; i < 100; i++ {
		result := IntN(n)
		if result < 0 || result >= n {
			t.Errorf("Expected integer in range [0,%d), got %d", n, result)
		}
	}
}

// TestIntRangeGeneration 测试区间整数生成功能
func TestIntRangeGeneration(t *testing.T) {
	min, max := 10, 20
	for i := 0; i < 100; i++ {
		result := IntRange(min, max)
		if result < min || result >= max {
			t.Errorf("Expected integer in range [%d,%d), got %d", min, max, result)
		}
	}
}

// TestFloat32Generation 测试浮点数生成功能
func TestFloat32Generation(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Float32()
		if result < 0.0 || result >= 1.0 {
			t.Errorf("Expected float32 in range [0.0,1.0), got %f", result)
		}
	}
}

// TestFloat64Generation 测试双精度浮点数生成功能
func TestFloat64Generation(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Float64()
		if result < 0.0 || result >= 1.0 {
			t.Errorf("Expected float64 in range [0.0,1.0), got %f", result)
		}
	}
}

// TestDifferentCharsets 测试不同字符集下的字符串生成
func TestDifferentCharsets(t *testing.T) {
	length := 10
	
	charsets := map[string]string{
		"Lowercase":      Lowercase,
		"Uppercase":      Uppercase,
		"Numeric":        Numeric,
		"Alphanumeric":   Alphanumeric,
		"Hex":            Hex,
		"Base64":         Base64,
		"URLSafeBase64":  URLSafeBase64,
	}
	
	for name, charset := range charsets {
		result := String(length, charset)
		if len(result) != length {
			t.Errorf("%s: Expected length %d, got %d", name, length, len(result))
		}
		
		for _, char := range result {
			if !strings.ContainsRune(charset, char) {
				t.Errorf("%s: Character %c not in charset %s", name, char, charset)
			}
		}
	}
}

// TestZeroLengthString 测试零长度字符串生成
func TestZeroLengthString(t *testing.T) {
	result := String(0, Alphanumeric)
	if result != "" {
		t.Errorf("Expected empty string for zero length, got %s", result)
	}
}

// TestNegativeLengthString 测试负长度字符串生成
func TestNegativeLengthString(t *testing.T) {
	result := String(-1, Alphanumeric)
	if result != "" {
		t.Errorf("Expected empty string for negative length, got %s", result)
	}
}

// BenchmarkString 测试String函数性能
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(100, Alphanumeric)
	}
}

// BenchmarkInt 测试Int函数性能
func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int()
	}
}

// BenchmarkIntN 测试IntN函数性能
func BenchmarkIntN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntN(1000)
	}
}