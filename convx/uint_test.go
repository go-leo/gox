package convx

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestToUnsignedE(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    interface{}
		f       func(any) (any, error)
		wantErr bool
	}{
		{
			name:  "int positive",
			input: 42,
			want:  uint(42),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "int negative",
			input: -1,
			want:  uint(0),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: true,
		},
		{
			name:  "uint positive",
			input: uint(42),
			want:  uint(42),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "int64 positive",
			input: int64(42),
			want:  uint64(42),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: false,
		},
		{
			name:  "int64 negative",
			input: int64(-1),
			want:  uint64(0),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: true,
		},
		{
			name:  "float64 positive",
			input: float64(42.0),
			want:  uint64(42),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: false,
		},
		{
			name:  "float64 negative",
			input: float64(-1.0),
			want:  uint64(0),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: true,
		},
		{
			name:  "string valid uint",
			input: "42",
			want:  uint64(42),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: false,
		},
		{
			name:  "string invalid uint",
			input: "abc",
			want:  uint64(0),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: true,
		},
		{
			name:  "bool true",
			input: true,
			want:  uint(1),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "bool false",
			input: false,
			want:  uint(0),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "nil",
			input: nil,
			want:  uint(0),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "time.Duration positive",
			input: 42 * time.Second,
			want:  uint64(42 * time.Second),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: false,
		},
		{
			name:  "time.Duration negative",
			input: -42 * time.Second,
			want:  uint64(0),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: true,
		},
		{
			name:  "time.Weekday valid",
			input: time.Sunday,
			want:  uint(0), // Assuming Sunday is represented as 0 in time.Weekday
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "time.Month valid",
			input: time.January,
			want:  uint(1), // Assuming January is represented as 1 in time.Month
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "json.Number positive",
			input: json.Number("42"),
			want:  uint64(42),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: false,
		},
		{
			name:  "json.Number negative",
			input: json.Number("-1"),
			want:  uint64(0),
			f: func(a any) (any, error) {
				return ToUint64E(a)
			},
			wantErr: true,
		},
		{
			name:  "driver.Valuer valid",
			input: sql.NullInt64{Int64: 42, Valid: true},
			want:  uint(42),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: false,
		},
		{
			name:  "driver.Valuer error",
			input: sql.NullInt64{Int64: -42, Valid: true},
			want:  uint(0),
			f: func(a any) (any, error) {
				return ToUintE(a)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("toUnsignedE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toUnsignedE() = %v, want %v", got, tt.want)
			}
		})
	}
}
