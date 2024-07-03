package convx

import (
	"fmt"
	"strconv"
	"testing"
)

func TestToStringerE(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "string input",
			input:   "hello",
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "bool input",
			input:   true,
			want:    "true",
			wantErr: false,
		},
		{
			name:    "float64 input",
			input:   3.14,
			want:    "3.14",
			wantErr: false,
		},
		{
			name:    "float32 input",
			input:   float32(3.14),
			want:    "3.14",
			wantErr: false,
		},
		{
			name:    "int input",
			input:   123,
			want:    "123",
			wantErr: false,
		},
		{
			name:    "int64 input",
			input:   int64(1234567890),
			want:    "1234567890",
			wantErr: false,
		},
		{
			name:    "int32 input",
			input:   int32(12345),
			want:    "12345",
			wantErr: false,
		},
		{
			name:    "int16 input",
			input:   int16(123),
			want:    "123",
			wantErr: false,
		},
		{
			name:    "int8 input",
			input:   int8(12),
			want:    "12",
			wantErr: false,
		},
		{
			name:    "uint input",
			input:   uint(123),
			want:    "123",
			wantErr: false,
		},
		{
			name:    "uint64 input",
			input:   uint64(1234567890),
			want:    "1234567890",
			wantErr: false,
		},
		{
			name:    "uint32 input",
			input:   uint32(12345),
			want:    "12345",
			wantErr: false,
		},
		{
			name:    "uint16 input",
			input:   uint16(123),
			want:    "123",
			wantErr: false,
		},
		{
			name:    "uint8 input",
			input:   uint8(12),
			want:    "12",
			wantErr: false,
		},
		{
			name:    "[]byte input",
			input:   []byte("hello"),
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "fmt.Stringer input",
			input:   strconv.FormatInt(42, 10),
			want:    "42",
			wantErr: false,
		},
		{
			name:    "error input",
			input:   fmt.Errorf("an error"),
			want:    "an error",
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    "",
			wantErr: false,
		},
		{
			name:    "unsupported type input",
			input:   struct{}{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringE(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("toStringerE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toStringerE() = %v, want %v", got, tt.want)
			}
		})
	}
}
