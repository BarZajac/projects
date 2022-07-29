package main

import "fmt"

func main() {
	fmt.Println("==== Slices declaration.")
	// Slice declaration.
	//var s0 []int
	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))

	s1 = append(s1, 4, 5)
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))

	s1 = append(s1, 4, 5, 6)
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))

	s3 := make([]int, 1, 20)
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3))

	s3 = append(s3, 1, 2, 3)
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3))
	fmt.Println(s3)

	s3 = s3[0:12]
	fmt.Println(s3)

	s3[11] = 11
	fmt.Println(s3)

	s3 = s3[:5]
	fmt.Println(s3)

	s3 = s3[:19]
	fmt.Println(s3)

	fmt.Println("==== Array declaration.")
	var a1 [2]int
	a2 := [2]int{0, 1}
	a3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(a1, a2)
	a2[1] = 10
	fmt.Println(a2)
	fmt.Println(a3)

	// Change array to slice.
	sa := a3[:]
	sa = append(sa, 100)
	fmt.Println(sa)
}
