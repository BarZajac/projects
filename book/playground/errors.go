package main

import (
	"errors"
	"fmt"
)

var errTooHigh = errors.New("")
var errTooSmall = errors.New("")

func main() {
	var i = -200
	a, err := errs(i)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, errTooSmall) {
			fmt.Println("Try bigger number")
		}
		if errors.Is(err, errTooHigh) {
			fmt.Println("Try smaller number")
		}
		return
	}
	fmt.Printf("Answer is %d", a)
}
func errs(a int) (int, error) {
	if a < 0 {
		return 0, errTooSmall

	}
	if a > 100 {
		return 0, errTooHigh
	}
	return a * 2, nil
}
