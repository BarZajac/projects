package main

import "fmt"

func main() {
	fmt.Println("==== String declaration.")
	// Declaration.
	var str0 string

	// Declare and assign.
	str1 := "abcdefghij"
	var str2 = "xyz"

	fmt.Printf("str0: '%s'\n", str0)
	fmt.Printf("str1: '%s'\n", str1)
	fmt.Printf("str2: '%s'\n", str2)

	fmt.Println("==== String indexing.")
	fmt.Printf("%d\n", str1[1])
	fmt.Printf("%s\n", str1[1:3])

	// Strings are immutable.
	//str1[2] = 'a'

	// Extra.
	// https://commons.wikimedia.org/wiki/File:ASCII-Table.svg
	buf0 := []byte(str1)
	fmt.Println(buf0)
	buf0[0] = 65
	fmt.Println(buf0)
	fmt.Println(string(buf0))
	buf0[1] = 'B'
	fmt.Println(string(buf0))
}
