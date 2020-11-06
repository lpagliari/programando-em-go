package main

import (
	"fmt"
)

func main() {
	PrintVersion()
	PrintHello("Luiza")
	PrintFullHello("Luiza", "Pagliari")
	PrintHelloAndAge("Luiza", "Pagliari", 38)

	fmt.Println("------------------")
	fmt.Printf("%d + %d = %d\n", 10, 7, Sum(10, 7))

	fmt.Println("------------------")
	price := 1000.0
	fmt.Printf("Price: %.2f\n", price)

	priceWithTax, priceInDolar := Convert(price)
	fmt.Printf("Price with tax %.2f\n", priceWithTax)
	fmt.Printf("Price in US$ %.2f\n", priceInDolar)

	priceWithTax, priceInDolar = Convert2(price)
	fmt.Printf("Price2 with tax %.2f\n", priceWithTax)
	fmt.Printf("Price2 in US$ %.2f\n", priceInDolar)

	fmt.Println("------------------")
	PrintFilePaths("/tmp", "one")
	PrintFilePaths("/tmp", "one", "two", "three")
}

// void, no parameter
func PrintVersion() {
	fmt.Println("1.0")
}

// void, with params
func PrintHello(name string) {
	fmt.Printf("Hello %s!\n", name)
}

// void, with params of same type
func PrintFullHello(name, surname string) {
	fmt.Printf("Hello %s %s!\n", name, surname)
}

// void, with more complex params
func PrintHelloAndAge(name, surname string, age int) {
	fmt.Printf("Hello %s %s! You are %d years old.\n", name, surname, age)
}

// with single return value
func Sum(x, y int) int {
	return x + y
}

// with multiple return values
func Convert(price float64) (float64, float64) {
	withTax := 1.0638
	toDolar := 5.5

	priceWithTax := price * withTax
	priceInDolar := price * toDolar

	return priceWithTax, priceInDolar
}

// with named return values
func Convert2(price float64) (priceWithTax, priceInDolar float64) {
	priceWithTax, priceInDolar = Convert(price) // don't need the `:=`, vars already exist
	return priceWithTax, priceInDolar
}

// with variable params
func PrintFilePaths(dir string, files ...string) {
	for _, file := range files {
		fmt.Printf("%s/%s.txt\n", dir, file)
	}
}
