package main

import (
	"fmt"
)
func main() {
	numbers := []int{2, 17, 43, 20, 18, 123}
	done := make(chan bool)
	odd, even := make(chan int), make(chan int)

	go split(numbers, odd, even, done)

	var oddNumbers, evenNumbers []int
	isDone := false
	for !isDone {
		select {
		case n := <- odd:
			oddNumbers = append(oddNumbers, n)
		case n := <- even:
			evenNumbers = append(evenNumbers, n)
		case isDone = <- done:
		}
	}

	fmt.Println("Odd numbers:", oddNumbers)
	fmt.Println("Even numbers:", evenNumbers)
}

func split(numbers []int, odd, even chan<- int, done chan<- bool) {
	for _, n := range numbers {
		if n % 2 == 0 {
			even <- n
		} else {
			odd <- n
		}
	}

	done <- true
}
