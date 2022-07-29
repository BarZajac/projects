package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	//ev := os.Environ()
	//for i, e := range ev {
	//	fmt.Println(i, e)
	//}
	//fmt.Println(os.Getenv("PWD"))

	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(sub(10, 1, 2, 3, 4))
	fmt.Println(sub(0, 1, 2, 3, 4))
}

func add(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func sub(initial int, a ...int) int {
	for _, v := range a {
		initial -= v
	}
	return initial
}

// env | grep PATH
// PATH_BARTEK=aaa
// env | grep PATH
// export PATH_BARTEK=aaa
// env | grep PATH

// cd /home/bartek/ws/go/lessons/env
// go run main.go
// env | grep PATH
// go run main.go
// go run main.go | grep PATH
// go build main.go
// ll
// ./main
// ./main | grep PATH
// PATH_BARTEK=bbb ./main | grep PATH
// env | grep PATH
// ./main
// go run main.go
// go run main.go -ggg -aaa
// go run main.go -ggg -aaa=1
// go run main.go --db-pass=1234
// cd ../inout/
// go run main.go
// go run main.go 1> 1.txt 2> 2.txt
// ll
// cat 1.txt
// cat 2.txt
// grep -R ala /etc
// grep -R ala /etc 2>2.txt
// grep -R ala /etc 2>2.txt | less
// cat 2.txt
// grep -R ala /etc 2>/dev/null | less
// grep -R ala /etc 2>/dev/null
// cat ~/.bash_history
