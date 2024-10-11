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
			a:    10,
			b:    20,
			want: 20,
		},
		{
			name: "Test 2",
			a:    20,
			b:    10,
			want: 20,
		},
		{
			name: "Test 3",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "Test 4",
			a:    -10,
			b:    -20,
			want: -10,
		},
		{
			name: "Test 5",
			a:    -20,
			b:    -10,
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxNum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("MaxNum() = %v, want %v", got, tt.want)
			}
			t.Logf("MaxNum(%v, %v) = %v", tt.a, tt.b, got)
		})
	}
}
