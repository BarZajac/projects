package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("==== Integer declaration and memory sizes.")

	// Declare integer.
	var a int
	var ua uint

	// Declare and assign.
	b := 2
	var c = 3

	// Unsigned integer types.
	var d uint8
	var e uint16
	var f uint32
	var g uint64

	// Signed integer types.
	var h int8
	var i int16
	var j int32
	var k int64

	// Integer sizes.
	fmt.Printf("a: %T, size: %d B, min: %d, max: %d\n", a, unsafe.Sizeof(a), math.MinInt, math.MaxInt)
	fmt.Printf("ua: %T, size: %d B, min: %d, max: %s\n", ua, unsafe.Sizeof(a), 0, strconv.FormatUint(math.MaxUint64, 10))
	fmt.Printf("b: %T, size: %d B, min: %d, max: %d\n", b, unsafe.Sizeof(b), math.MinInt, math.MaxInt)
	fmt.Printf("c: %T, size: %d B, min: %d, max: %d\n", c, unsafe.Sizeof(c), math.MinInt, math.MaxInt)
	fmt.Println()

	fmt.Printf("d: %T, size: %d B, min: %d, max: %d\n", d, unsafe.Sizeof(d), 0, math.MaxUint8)
	fmt.Printf("e: %T, size: %d B, min: %d, max: %d\n", e, unsafe.Sizeof(e), 0, math.MaxUint16)
	fmt.Printf("f: %T, size: %d B, min: %d, max: %d\n", f, unsafe.Sizeof(f), 0, math.MaxUint32)
	fmt.Printf("g: %T, size: %d B, min: %d, max: %s\n", g, unsafe.Sizeof(g), 0, strconv.FormatUint(math.MaxUint64, 10))
	fmt.Println()

	fmt.Printf("h: %T, size: %d B, min: %d, max: %d\n", h, unsafe.Sizeof(h), math.MinInt8, math.MaxInt8)
	fmt.Printf("i: %T, size: %d B, min: %d, max: %d\n", i, unsafe.Sizeof(i), math.MinInt16, math.MaxInt16)
	fmt.Printf("j: %T, size: %d B, min: %d, max: %d\n", j, unsafe.Sizeof(j), math.MinInt32, math.MaxInt32)
	fmt.Printf("k: %T, size: %d B, min: %d, max: %d\n", k, unsafe.Sizeof(k), math.MinInt64, math.MaxInt64)
	fmt.Println()
}
