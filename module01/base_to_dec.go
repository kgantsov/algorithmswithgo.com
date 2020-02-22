package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	letters := "0123456789ABCDEF"
	value = Reverse(value)
	var res int
	for i, el := range value {
		n := strings.Index(letters, string(el))

		res = res + n*int(math.Pow(float64(base), float64(i)))

	}
	return res
}
