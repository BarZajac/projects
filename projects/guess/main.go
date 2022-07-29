package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var ErrEnd = errors.New("end of input")

func main() {
	fmt.Println("Guess a number from 0 to 100. Good luck!")
	guess := rand.Intn(101)

	for {
		fmt.Print("Enter text: ")
		n, err := ReadNumber()
		if err != nil {
			if errors.Is(err, ErrEnd) {
				fmt.Println("Bye!")
				return
			}
			if errors.Is(err, strconv.ErrSyntax) {
				fmt.Println("Not a number!")
				return
			}
			fmt.Println(err)
			return
			// continue
			// break
		}

		switch {
		case n < guess:
			fmt.Println("Your guess is to low.")
		case n > guess:
			fmt.Println("Your guess is to high.")
		default:
			fmt.Println("You guessed it!")
		}

		fmt.Println("Try again or press ENTER to exit.")
	}
}

func ReadNumber() (int, error) {
	rdr := bufio.NewReader(os.Stdin)

	// ReadString will block until the delimiter is entered.
	in, err := rdr.ReadString('\n')
	if err != nil {
		return 0, errors.New("an error occurred while reading input. " +
			"Please try again")
	}

	if in == "\n" {
		return 0, ErrEnd
	}

	// Remove the delimiter from the string.
	in = strings.TrimSuffix(in, "\n")

	n, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}

	return n, nil
}
