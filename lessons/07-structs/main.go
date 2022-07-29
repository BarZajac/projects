package main

import "fmt"

// Player defines custom type with two fields.
type Player struct {
	Name  string // Name field.
	Score int    // Score Field.
}

func main() {
	//i := 5
	//
	//println(i)
	//println(&i)
	//addToInt(&i)
	//println(i)
	//
	//return

	var p1, p2 Player

	p1.Name = "Bartek"
	p1.Score = 110
	ex(p1)

	p2.Name = "Rafal"
	p2.Score = 101
	ex(p2)

	fmt.Println()

	//p2 = addToScore(p2, 10)
	addToScorePtr(&p2, 10)
	score(p1, p2)
}

func ex(p Player) {
	fmt.Printf("Palyer: %s, Score: %d\n", p.Name, p.Score)
}

func score(player1, player2 Player) {
	if player1.Score > player2.Score {
		ex(player1)
	} else {
		ex(player2)
	}
}

func addToScore(p Player, v int) Player {
	p.Score = p.Score + v
	return p
	//p.Score += v
}

func addToScorePtr(p *Player, v int) {
	p.Score = p.Score + v
}

func addToInt(v *int) {
	println(v)
	println(&v)
	*v = 10
}
