package chanx

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// TestAsSlice tests the AsSlice function for a slice of integers.
func TestAsSlice(t *testing.T) {
	// Create a channel and send some values
	ch := make(chan int, 3) // Buffered channel to avoid blocking
	ch <- 1
	ch <- 2
	ch <- 3

	// Close the channel to end the range loop in AsSlice
	close(ch)

	// Call the AsSlice function
	result := AsSlice(context.Background(), ch)

	// Define the expected result
	expected := []int{1, 2, 3}

	// Use reflect.DeepEqual to compare the slices
	if !reflect.DeepEqual(<-result, expected) {
		t.Errorf("AsSlice() = %v, want %v", result, expected)
	}
}

// TestAsSliceEmpty tests the AsSlice function with an empty channel.
func TestAsSliceEmpty(t *testing.T) {
	// Create an empty channel
	ch := make(chan int)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	// Call the AsSlice function
	result := AsSlice(ctx, ch)

	// Define the expected result for an empty channel
	expected := []int{}

	// Use reflect.DeepEqual to compare the slices
	if !reflect.DeepEqual(<-result, expected) {
		t.Errorf("AsSlice() = %v, want %v", result, expected)
	}
}

func TestAsyncCall(t *testing.T) {
	var funcs []func()
	for i := 0; i < 100; i++ {
		n := i
		funcs = append(funcs, func() { fmt.Println(n) })
	}

	for _, f := range funcs {
		go f()
	}

	<-time.After(100 * time.Second)
}
