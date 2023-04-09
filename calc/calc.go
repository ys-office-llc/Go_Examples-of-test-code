// Package calc provides basic arithmetic operations.
package calc

// Max returns the maximum of two integers x and y.
// If x and y are equal, x is returned.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
