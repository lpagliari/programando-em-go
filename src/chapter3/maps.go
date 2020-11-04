package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	statistics := collectStatistics(words)
	printStatistics(statistics)
}

func collectStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		char := strings.ToUpper(string(word[0]))


		// here's different from the book: `counter` is 0 when it does not exist, so we can simplify the
		// if/else code above (from the book) and use a single line to update the counter
		counter, _ := statistics[char]
		statistics[char] = counter + 1

		// original code from the book:
		// counter, exists := statistics[char]
		// if exists {
		// 	statistics[char] = counter + 1
		// } else {
		// 	statistics[char] = 1
		// }
	}
	return statistics
}

func printStatistics(statistics map[string]int) {
	fmt.Println("Counter of words, by their initial char:")

	for char, counter := range statistics {
		fmt.Printf("%s: %d\n", char, counter)
	}
}
