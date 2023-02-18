package shanks

import (
	"fmt"

	"github.com/jabbalaci/shanksbot-go/lib/jassert"
)

// const SHOW_DIGITS = true

const SHOW_DIGITS = false

func Shanks(n int) int {
	jassert.Assert(n > 0, "n must be positive")
	//
	if n == 1 {
		return 0
	}
	//
	result := 0

	cache := make([]int, 10+1)
	for i := range cache {
		cache[i] = i * n
	}
	found := make(map[int]bool)

	number := 1
	for {
		if number == 0 {
			break
		}
		number *= 10
		if found[number] {
			break
		}
		// else
		found[number] = true
		for i, curr := range cache {
			if curr > number {
				result++
				if SHOW_DIGITS {
					fmt.Print(i - 1)
				}
				number = number - cache[i-1]
				break
			}
		}
	}
	if SHOW_DIGITS {
		fmt.Println()
	}
	return result
}
