package convx

import (
	"reflect"
	"testing"
)

func TestToSignedE(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		want      any
		wantErr   bool
		err       error
		toSignedE func(any) (any, error)
	}{
		{
			name:    "int value",
			input:   10,
			want:    10,
			wantErr: false,
			toSignedE: func(a any) (any, error) {
				return ToSignedE[int](a)
			},
		},
		{
			name:    "int64 value",
			input:   int64(10),
			want:    int64(10),
			wantErr: false,
			toSignedE: func(a any) (any, error) {
				return ToSignedE[int64](a)
			},
		},
		{
			name:    "int32 value",
			input:   int32(10),
			want:    int32(10),
			wantErr: false,
			toSignedE: func(a any) (any, error) {
				return ToSignedE[int32](a)
			},
		},
		// Add more test cases for other types
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.toSignedE(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSignedE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("ToSignedE() error = %v, want %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("ToSignedE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSignedSliceE(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    []int
		wantErr bool
	}{
		{
			name:    "nil input",
			input:   nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "slice input",
			input:   []int{1, 2, 3},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "array input",
			input:   [3]int{4, 5, 6},
			want:    []int{4, 5, 6},
			wantErr: false,
		},
		{
			name:    "wrong type input",
			input:   "abc",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSignedSliceE[[]int](tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSignedSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSignedSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}
