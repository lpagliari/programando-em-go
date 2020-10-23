package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)

		if err != nil {
			fmt.Printf("Entry %s is not a valid integer!\n", n)
			os.Exit(1)
		}

		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	size := len(numbers)

	if (size <= 1) {
		return numbers
	}

	pivotIndex := size / 2
	pivot := numbers[pivotIndex]
	// here's different from the book: there's no need to create a copy of `numbers`, as we won't
	// be changing anything on that slice. So we can simply create a new `numbersWithoutPivot` calling
	// only `append`
	numbersWithoutPivot := append(numbers[: pivotIndex], numbers[pivotIndex + 1 :]...)

	lessThanPivot, greaterThanPivot := partitionate(numbersWithoutPivot, pivot)

	l := quicksort(lessThanPivot)
	g := quicksort(greaterThanPivot)
	return append(append(l, pivot), g...)
}

func partitionate(numbers []int, pivot int) (lessThanPivot, greaterThanPivot []int) {
	for _, number := range numbers {
		if number <= pivot {
			lessThanPivot = append(lessThanPivot, number)
		} else {
			greaterThanPivot = append(greaterThanPivot, number)
		}
	}
	return lessThanPivot, greaterThanPivot
}
