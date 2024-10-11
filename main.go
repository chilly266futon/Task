package myMath

import (
	"math"
)

// MaxNum returns the maximum of two float64 numbers.
// It uses the formula (a + b + |a-b|) / 2 to determine the maximum value.
//
// Parameters:
//   - a: the first float64 number.
//   - b: the second float64 number.
//
// Returns:
//   - The maximum of the two float64 numbers.
func MaxNum(a, b float64) float64 {
	res := (a + b + math.Abs(a-b)) / 2
	return res
}
