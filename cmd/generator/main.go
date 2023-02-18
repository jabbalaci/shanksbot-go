package main

import (
	"fmt"

	"github.com/jabbalaci/shanksbot-go/lib/jmath"
	"github.com/jabbalaci/shanksbot-go/lib/shanks"
)

func main() {
	n := 110_000
	primes := jmath.GetPrimesBelow(n)

	for _, p := range primes {
		fmt.Printf("%v -> %v\n", p, shanks.Shanks(p))
	}
}
