package main

import (
	"bytes"
	"github.com/go-leo/gox/encodingx/jsonx"
)

func main() {
	buf := &bytes.Buffer{}
	encoder := jsonx.NewEncoder(buf)
	err := encoder.Encode(map[string]string{"hello": "world"})
	if err != nil {
		panic(err)
	}

}
