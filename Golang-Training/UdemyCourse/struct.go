// Structs are a collection of fields.
// They're useful for grouping data togheter to form records.
// the closest thing in Go to objects/OOP
package main

import "fmt"

type player struct {
	firstname string
	lastname  string
	score     int
	//any       bool
}

type game struct {
	players map[string]*player
}

func (g *game) getWinner() *player {
	var maxScore int
	var winner *player

	for _, value := range g.players {
		if value.score > maxScore {
			winner = value
			maxScore = value.score
		}
	}
	return winner
}

func main() {

	g := &game{
		players: make(map[string]*player),
	}

	p1 := player{}
	fmt.Println("p1", p1)

	g.players["p1"] = &p1
	fmt.Println("game", g)

	var p2 player

	p2 = player{}
	p2.score = 4
	fmt.Println("p2", p2)

	g.players["p2"] = &p2
	fmt.Println("game", g)

	p3 := player{
		lastname:  "Cortes",
		firstname: "Manuel",
	}
	fmt.Println("p3", p3)
	p3.score = 10
	fmt.Println("p3", p3)

	fmt.Println("g.p[2]", g.players["p2"])
	fmt.Println("g.p[2]", *g.players["p2"])

	winner := g.getWinner()
	fmt.Println("winner", winner)
	fmt.Println("*winner", *winner)
}
