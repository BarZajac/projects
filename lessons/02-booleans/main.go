package main

import "fmt"

func main() {
	// Booleans represent true or false values.
	// var b bool

	b := false
	if b {
		fmt.Println("b is true")
	} else {
		fmt.Println("b is false")
	}

	var tf bool
	i := 11
	j := 90
	tf = i > 10 && j < 10 // true && false -> false
	fmt.Println("value: ", tf)
	fmt.Println("i == i: ", i == i)
	fmt.Println("i == j: ", i == j)
	fmt.Println("i < j: ", i < j)

	if i > 10 && j < 10 {
		fmt.Println("value is bigger then 10")
	} else {
		fmt.Println("value is smaller then 10")
	}

	i = 10
	if i > 10 {
		fmt.Println("A")
	} else if i < 10 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	switch {
	case i > 10:
		fmt.Println("A")
	case i < 10:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	switch i {
	case 10:
		fmt.Println("A")
	case 11:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}
