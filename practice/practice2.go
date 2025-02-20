package main

import (
	"errors"
	"fmt"
)

func printMe(value string) {
	fmt.Println(value) // Print the value passed as an argument
}

func main() {
	var printValue string = "Hello World"
	printMe(printValue) // Call printMe with the value "Hello World"

	var numerator int = 11
	var denominator int = 2
	var result int
	var err error

	result, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(result) // Output the result of integer division
}

// Define the intDivision function

func intDivision(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("Can't divide by Zero")
	}

	var result int = numerator / denominator
	return result, nil
}
