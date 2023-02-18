// Package jmath faciliates work with numbers.
package jmath

import (
	"math"
)

// Returns all prime numbers below the given threshold.
// The upper limit `n` is not included.
// It uses the sieve of Eratosthenes algorithm.
func GetPrimesBelow(n int) []int {
	result := []int{}
	lst := make([]bool, n)
	for i := range lst {
		lst[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if lst[i] {
			for j := i * i; j < n; j += i {
				lst[j] = false
			}
		}
	}
	for i := 2; i < n; i++ {
		if lst[i] {
			result = append(result, i)
		}
	}

	return result
}
