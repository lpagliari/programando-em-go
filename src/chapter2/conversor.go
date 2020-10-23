package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsLength := len(os.Args)
	lastArg := argsLength - 1

	if argsLength < 3 {
		fmt.Println("usage: conversor <values> <unit>")
		fmt.Println("\nUnit: \tcelsius | kilometers")
		os.Exit(1)
	}

	sourceUnit := os.Args[lastArg]
	sourceValues := os.Args[1 : lastArg]

	var targetUnit string
	if sourceUnit == "celsius" {
		targetUnit = "fahrenheit"
	} else if sourceUnit == "kilometers" {
		targetUnit = "miles"
	} else {
		fmt.Printf("Unknown unit \"%s\". Expected one of: celsius | kilometers.\n", sourceUnit)
		os.Exit(1)
	}

	for i, v := range sourceValues {
		sourceValue, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("Value %s on position %d not valid!\n", v, i)
			os.Exit(1)
		}

		var targetValue float64
		if sourceUnit == "celsius" {
			targetValue = 1.8 * sourceValue + 32
		} else if sourceUnit == "kilometers" {
			targetValue = sourceValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", sourceValue, sourceUnit, targetValue, targetUnit)
	}
}
