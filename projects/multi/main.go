package main

import (
	"fmt"
)

func main() {
	//var i int
	//fmt.Println("  1 2 3 4 5 6 7 8 9 10")
	//for i = 1; i <= 10; i++ {
	//	i2 := i * 2
	//	i3 := i * 3
	//	i4 := i * 4
	//	i5 := i * 5
	//	i6 := i * 6
	//	i7 := i * 7
	//	i8 := i * 8
	//	i9 := i * 9
	//	i10 := i * 10
	//	fmt.Println(i, i, i2, i3, i4, i5, i6, i7, i8, i9, i10)
	//}

	for i := 0; i <= 10; i++ {
		if i == 0 {
			fmt.Printf("     ")
			continue
		}
		fmt.Printf(" %4d", i)
	}
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %4d", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf(" %4d", i*j)
		}
		fmt.Println()
	}

}
func nul(n int) {

}
