package main

import "fmt"

func main() {
	//i := 5
	//p := &i
	//fmt.Printf("main: %d, %p\n", i, p)
	//do(i)
	//fmt.Printf("main: %d, %p\n", i, p)
	//
	//fmt.Println("----------------")
	//
	//doPointer(p)
	//fmt.Printf("main: %d, %p\n", i, p)

	//a := []int{0, 1, 2}
	//sliceChange(a)
	//fmt.Println(a)

	pl := Player{
		Name:  "Bartek",
		Value: 1,
	}
	changeMyType(pl)
	fmt.Println(pl)

	changeMyTypeByPointer(&pl)
	fmt.Println(pl)
}

func do(i int) {
	i *= 5
	//i = i * 5
	fmt.Printf("in functin: %d, %p\n", i, &i)
}

func doPointer(p *int) {
	*p = 100
	fmt.Printf("in functin: %d, %p\n", *p, p)
}

func sliceChange(a []int) {
	a[1] = 100
}

func changeMyType(player Player) { // Przez wartosc / By Value
	player.Value = 100
}

func changeMyTypeByPointer(player *Player) { // Przez wskaznik / By pointer
	player.Value = 100
}

type Player struct {
	Name  string
	Value int
}

func fnP(g int, a []*func(int, string) error) bool {
	switch g {
	case 3:
		if err := (*a[3])(1, "bartek"); err != nil {
			return true
		}
	}
	return false
}
