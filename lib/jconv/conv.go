package jconv

import (
	"strconv"
)

// Convert a string to int.
// Example: "42" (string) -> 42 (int)
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
