package errorx

import (
	"errors"
	"testing"
)

const expectedIntValue = 43

// 假设这是你要测试的原函数
func TestBreak(t *testing.T) {
	// 场景1: pre为nil
	{
		var preErr error
		result, err := Break[int](preErr)(testFuncThatReturnsIntAndError)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expectedIntValue { // 假设你有一个预期的int值
			t.Errorf("Expected result %d, got %d", expectedIntValue, result)
		}
	}

	// 场景2: pre不为nil
	{
		preErr := errors.New("Test error")
		result, err := Break[int](preErr)(testFuncThatReturnsIntAndError)

		if err == nil || err.Error() != preErr.Error() {
			t.Errorf("Expected error '%s', got '%v'", preErr, err)
		}
		if result != 0 { // T类型的零值
			t.Errorf("Expected zero value for result, got %d", result)
		}
	}
}

// 这是一个用于测试的辅助函数，模拟返回一个int和一个error
func testFuncThatReturnsIntAndError() (int, error) {
	return expectedIntValue, nil // 返回预期的int值和nil错误
}

// 假设这是你要测试的原函数
func TestFunction(t *testing.T) {
	// 场景1: f成功，pre不为nil
	{
		preErr := errors.New("Pre-error")
		fResult, _ := testFuncThatReturnsIntAndError()
		result, err := Continue[int](preErr)(testFuncThatReturnsIntAndError)
		if err == nil || err.Error() != preErr.Error() {
			t.Errorf("Expected error '%s', got '%v'", preErr, err)
		}
		if result != fResult {
			t.Errorf("Expected result %d, got %d", expectedIntValue, result)
		}
	}

	// 场景2: f成功，有错误，pre为nil
	{
		var preErr error
		fResult, fErr := testFuncWithErr()
		result, err := Continue[int](preErr)(testFuncWithErr)
		if err == nil {
			t.Error("Expected an error, got none")
		}
		if err.Error() != fErr.Error() {
			t.Errorf("Expected error '%s', got '%v'", fErr, err)
		}
		if result != fResult {
			t.Errorf("Expected result %d, got %d", fResult, result)
		}
	}

	// 场景3: f失败，有错误，pre为nil
	{
		var preErr error
		fResult, fErr := testFuncWithErr()
		result, err := Continue[int](preErr)(testFuncWithErr)

		if err == nil {
			t.Error("Expected an error, got none")
		}
		if err.Error() != fErr.Error() {
			t.Errorf("Expected error '%s', got '%v'", fErr, err)
		}
		if result != fResult {
			t.Errorf("Expected result %d, got %d", fResult, result)
		}
	}

	// 场景4: f失败，有错误，pre不为nil
	{
		preErr := errors.New("Pre-error")
		fResult, fErr := testFuncWithErr()
		result, err := Continue[int](preErr)(testFuncWithErr)

		combinedErrMsg := Join(preErr, fErr)
		if err == nil || err.Error() != combinedErrMsg.Error() {
			t.Errorf("Expected error '%s', got '%v'", combinedErrMsg, err)
		}
		if result != fResult {
			t.Errorf("Expected result %d, got %d", fResult, result)
		}
	}
}

// 这是一个用于测试的辅助函数，模拟返回一个int和一个错误
func testFuncWithErr() (int, error) {
	return 0, errors.New("Test error from f")
}
