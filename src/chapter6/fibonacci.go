package main

import (
	"fmt"
)

func main() {
	a, b := 0, 1

	fibonacci := func() int {
		newA := b
		newB := a+b

		a, b = newA, newB

		return a
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", fibonacci())
	}
}
