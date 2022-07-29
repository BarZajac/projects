package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 1
	b := 1.1

	sum, err := addPositive(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("sum = %f\n", sum)
	}
}

// example has no arguments and returns no values.
func example() {
	// no return
}

// withReturn returns integer.
func withReturn() int {
	return 1
}

func multipleReturn() (int, string) {
	return 2, "abc"
}

func addPositive(a int, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("must be positive")
	}
	return float64(a) + b, nil
}
