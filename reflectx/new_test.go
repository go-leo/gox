package reflectx

import (
	"reflect"
	"testing"
	"unsafe"
)

// TestNew is a test function for the New function.
func TestNew(t *testing.T) {
	i := New[int]()
	t.Log(i)
	ip := New[*int]()
	t.Log(ip)
	ipp := New[**int]()
	t.Log(ipp)

}

func Test_new(t *testing.T) {
	tests := []struct {
		name      string
		input     reflect.Type
		wantPanic bool
		want      reflect.Value
	}{
		{
			name:      "Int Ptr",
			input:     reflect.TypeOf((*int)(nil)),
			wantPanic: false,
			want:      reflect.ValueOf(new(int)),
		},
		{
			name:      "Bool",
			input:     reflect.TypeOf(true),
			wantPanic: false,
			want:      reflect.ValueOf(false),
		},
		{
			name:      "Int",
			input:     reflect.TypeOf(0),
			wantPanic: false,
			want:      reflect.ValueOf(0),
		},
		{
			name:      "Uint",
			input:     reflect.TypeOf(uint(0)),
			wantPanic: false,
			want:      reflect.ValueOf(uint(0)),
		},
		{
			name:      "Float64",
			input:     reflect.TypeOf(float64(0)),
			wantPanic: false,
			want:      reflect.ValueOf(float64(0)),
		},
		{
			name:      "Complex64",
			input:     reflect.TypeOf(complex64(0)),
			wantPanic: false,
			want:      reflect.ValueOf(complex64(0)),
		},
		{
			name:      "Array",
			input:     reflect.TypeOf([5]int{}),
			wantPanic: false,
			want:      reflect.ValueOf([5]int{}),
		},
		{
			name:      "Chan",
			input:     reflect.TypeOf(make(chan int)),
			wantPanic: false,
			want:      reflect.ValueOf(make(chan int)),
		},
		{
			name:      "Func",
			input:     reflect.TypeOf(func() {}),
			wantPanic: false,
			want:      reflect.ValueOf(func() {}),
		},
		{
			name:      "Interface",
			input:     reflect.TypeOf((*interface{})(nil)),
			wantPanic: false,
			want:      reflect.ValueOf((*interface{})(nil)),
		},
		{
			name:      "Map",
			input:     reflect.TypeOf(map[string]int{}),
			wantPanic: false,
			want:      reflect.ValueOf(map[string]int{}),
		},
		{
			name:      "Ptr",
			input:     reflect.TypeOf(new(int)),
			wantPanic: false,
			want:      reflect.ValueOf(new(int)),
		},
		{
			name:      "Slice",
			input:     reflect.TypeOf([]int{}),
			wantPanic: false,
			want:      reflect.ValueOf([]int{}),
		},
		{
			name:      "String",
			input:     reflect.TypeOf(""),
			wantPanic: false,
			want:      reflect.ValueOf(""),
		},
		{
			name:      "Struct",
			input:     reflect.TypeOf(struct{}{}),
			wantPanic: false,
			want:      reflect.ValueOf(struct{}{}),
		},
		{
			name:      "UnsafePointer",
			input:     reflect.TypeOf((*unsafe.Pointer)(nil)),
			wantPanic: false,
			want:      reflect.ValueOf((*unsafe.Pointer)(nil)),
		},
		{
			name:      "Unhandled case",
			input:     reflect.TypeOf((**int)(nil)),
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			got := reflectNew(tt.input)
			if !reflect.DeepEqual(got.Type(), tt.want.Type()) {
				t.Errorf("reflectNew() = %v, want %v", got.Type(), tt.want.Type())
			}
		})
	}
}
