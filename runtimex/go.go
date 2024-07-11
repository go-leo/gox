package runtimex

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GoID is used to retrieve the ID of the current Goroutine.
// It achieves this by capturing the stack information of the current goroutine and parsing the goroutine ID from it.
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	return parseGoID(string(buf[:n]))
}

func parseGoID(str string) int {
	ifField := strings.Fields(strings.TrimPrefix(str, "goroutine "))[0]
	id, err := strconv.Atoi(ifField)
	if err != nil {
		panic(fmt.Errorf("syncx: failed to get goroutine id, %w", err))
	}
	return id
}
