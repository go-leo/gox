package yijing

import "testing"


func Test十神_String(t *testing.T) {
    // Create a 十神 instance with a test name
    ss := 十神{名: "TestShishen"}

    // Call the String method and check if it returns the correct name
    result := ss.String()
    if result != "TestShishen" {
        t.Errorf("String() = %v, want %v", result, "TestShishen")
    }
}