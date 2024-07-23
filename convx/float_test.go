package convx

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestToFloatE(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		want      any
		wantErr   bool
		errTarget error
		toFloat   func(any) (any, error)
	}{
		{
			name:      "int value",
			input:     42,
			want:      42.0,
			wantErr:   false,
			errTarget: nil,
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		{
			name:      "int64 value",
			input:     int64(42),
			want:      42.0,
			wantErr:   false,
			errTarget: nil,
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		{
			name:      "float64 value",
			input:     float64(42.123),
			want:      42.123,
			wantErr:   false,
			errTarget: nil,
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		{
			name:      "string value",
			input:     "42.123",
			want:      42.123,
			wantErr:   false,
			errTarget: nil,
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		{
			name:      "invalid string value",
			input:     "invalid",
			want:      0.0,
			wantErr:   true,
			errTarget: errors.New(failedCast),
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		{
			name:      "bool true value",
			input:     true,
			want:      1.0,
			wantErr:   false,
			errTarget: nil,
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		{
			name:      "bool false value",
			input:     false,
			want:      0.0,
			wantErr:   false,
			errTarget: nil,
			toFloat: func(i any) (any, error) {
				return ToFloatE[float64](i)
			},
		},
		// Add more test cases for other types
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.toFloat(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				assert.True(t, strings.Contains(tt.errTarget.Error(), failedCast))
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
