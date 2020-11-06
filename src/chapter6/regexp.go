package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// \b: at ASCII word boundary
	// \w: a word character: [0-9A-Za-z_]
	exprToCaptureFirstCharOfWords := regexp.MustCompile("\\b\\w")

	transform := func(s string) string {
		return strings.ToUpper(s)
	}

	text := "luiza pagliari"

	fmt.Println(transform(text))
	fmt.Println(exprToCaptureFirstCharOfWords.ReplaceAllStringFunc(text, transform))
}
