package test

import (
	. "awesomeProject"
	"testing"
)

func Test_MaxNum(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{
			name: "Test 1",
			a:    57.0,
			b:    9.0,
			want: 57.0,
		},
		{
			name: "Test 2",
			a:    0.0,
			b:    0.0,
			want: 0.0,
		},
		{
			name: "Test 3",
			a:    -1.0,
			b:    -2.0,
			want: -1.0,
		},
		{
			name: "Test 4",
			a:    17.0,
			b:    23.0,
			want: 23.0,
		},
		{
			name: "Test 5",
			a:    0.0,
			b:    1.0,
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxNum(tt.a, tt.b); got != tt.want {
				t.Errorf("MaxNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
