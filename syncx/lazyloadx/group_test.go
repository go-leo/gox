package lazyloadx

import (
	"golang.org/x/sync/singleflight"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 定义一个测试用例结构体
type testCase struct {
	name       string
	group      *Group[string]
	key        string
	newFunc    func(key string) (string, error)
	want       string
	wantErr    error
	wantLoaded bool
}

func TestGroup_LoadOrNew(t *testing.T) {
	// 测试用例
	tests := []testCase{
		{
			name:       "Load existing value",
			group:      &Group[string]{m: sync.Map{}, g: singleflight.Group{}, New: func(key string) (string, error) { return "testValue", nil }},
			key:        "testKey",
			newFunc:    nil,
			want:       "testValue",
			wantErr:    nil,
			wantLoaded: false,
		},
		{
			name:       "Create new value",
			group:      &Group[string]{m: sync.Map{}, g: singleflight.Group{}},
			key:        "newKey",
			newFunc:    func(key string) (string, error) { return "newValue", nil },
			want:       "newValue",
			wantErr:    nil,
			wantLoaded: false,
		},
		{
			name:       "New function is nil",
			group:      &Group[string]{m: sync.Map{}, g: singleflight.Group{}},
			key:        "newKey",
			newFunc:    nil,
			want:       "",
			wantErr:    ErrNilFunction,
			wantLoaded: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the LoadOrNew method
			got, err, loaded := tt.group.LoadOrNew(tt.key, tt.newFunc)

			// Assertions
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantLoaded, loaded)

			// Assertions for the sync.Map
			if _, ok := tt.group.m.Load(tt.key); ok {
				tt.group.m.Delete(tt.key) // Clean up after the test
			}
		})
	}
}
