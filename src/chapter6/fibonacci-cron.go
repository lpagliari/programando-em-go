package main

import (
	"fmt"
	"time"
)

func main() {
	Cron(GenerateFibonacci(8))
	Cron(GenerateFibonacci(48))
	Cron(GenerateFibonacci(88))
}

func GenerateFibonacci(n int) func() {
	return func() {
			a, b := 0, 1

		fibonacci := func() int {
			newA := b
			newB := a+b

			a, b = newA, newB

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fibonacci())
		}
	}
}

func Cron(fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)

	fmt.Printf("\nTime to run: %s\n\n", duration)
}
