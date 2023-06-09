// compute a safe prime and its field generator
//
// Usage: ./primefield bits [bits ..]
//
// Copyright 2013-2023 arag0re <arag0re.eth-at-protonmail-dot-com>
// License: MIT
//

package main

import (
	"fmt"
	"github.com/arag0re/go-apple-srp6"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s bits [bits ..]\n", os.Args[0])
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		ss := os.Args[i]
		n, err := strconv.Atoi(ss)
		if err != nil || n <= 0 {
			fmt.Fprintf(os.Stderr, "%s: %s not an integer?\n", os.Args[0], ss)
			continue
		}

		p, g, err := srp.NewPrimeField(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: Can't generate %d bit prime field\n", os.Args[0], n)
			continue
		}

		fmt.Printf("%d:%d:%x\n", n, g, p)
	}
}
