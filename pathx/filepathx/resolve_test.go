package filepathx

import "testing"

func TestResolve(t *testing.T) {
	resolve, err := Resolve("/Users/stuff/Workspace/github/go-leo/cqrs/example/api/demo/demo.proto",
		"../../../internal/demo/query")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resolve)
}
