package main

import (
	"fmt"
	"os"

	"github.com/jabbalaci/shanksbot-go/lib/jconsole"
	"github.com/jabbalaci/shanksbot-go/lib/jconv"
	"github.com/jabbalaci/shanksbot-go/lib/shanks"
)

func main() {
	text := jconsole.Input("Number: ")
	n, err := jconv.StrToInt(text)
	if n <= 0 || err != nil {
		fmt.Fprintln(os.Stderr, "Error: provide a positive number")
		os.Exit(1)
	}

	result := shanks.Shanks(n)
	fmt.Println("Result:", result)
}
