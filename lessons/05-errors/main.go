package main

import (
	"errors"
	"fmt"
)

var ErrNegative = errors.New("value is negative")
var ErrTooHigh = errors.New("value is too high")

func main() {
	var i = 1
	d, err := doubleSmall(i)
	if err != nil {
		fmt.Println(err)

		if errors.Is(err, ErrNegative) {
			fmt.Println("try bigger number")
		}

		if errors.Is(err, ErrTooHigh) {
			fmt.Println("try smaller number")
		}

		return
	}

	fmt.Printf("answer %d\n", d)
}

func doubleSmall(a int) (int, error) {
	if a > 100 {
		return 0, ErrTooHigh
	}

	if a < 0 {
		return 0, ErrNegative
	}

	return 2 * a, nil
}
