package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	scn := bufio.NewScanner(r)

	for scn.Scan() {
		lin := scn.Text()
		lin = strings.ToUpper(lin)
		fmt.Println(lin)
	}
	if err := scn.Err(); err != nil {
		_, _ = io.WriteString(os.Stderr, err.Error())
	}
}
