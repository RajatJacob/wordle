package main

import (
	"github.com/RajatJacob/wordle/game"
)

func main() {
	g := game.Create()
	g.Guess("BLACK", [game.WORD_LEN]int{0, 0, 2, 0, 0})
	g.Guess("SMART", [game.WORD_LEN]int{0, 1, 2, 1, 0})
	// g.Guess("MEMES", [game.WORD_LEN]int{0, 1, 0, 0, 2})
	// g.Guess("ELDER", [game.WORD_LEN]int{2, 2, 2, 2, 2})
	g.Print()
}
