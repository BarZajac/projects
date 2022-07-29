package main

import (
	"fmt"
)

func main() {
	var a = 50.0
	var b = a*5 - 120
	var c = b - a
	var d = 3.6 * c
	fmt.Printf("Samolot lecial z predkocia %f metrow na sekunde.\nJest to %f kilometrow na godzine ", c, d)
}

//func main() {
//	var a = 50.0
//	var b = a*5 - 120
//	var c = b - a
//	var d = 3.6
//	var e = d * c
//	fmt.Printf("Samolot lecial z predkocia %f metrow na sekunde.\nJest to %f kilometrow na godzine ", c, e)

//}
