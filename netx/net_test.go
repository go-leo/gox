package netx

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestEmpty(t *testing.T) {
	parse, err := url.Parse("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parse)
}

func TestServer(t *testing.T) {
	t.Log(net.JoinHostPort("", strconv.Itoa(0)))
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})
	listen, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(0)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(listen.Addr().String())
	http.Serve(listen, mux)

}
