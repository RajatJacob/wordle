package main

import (
	"github.com/RajatJacob/wordle/game"
)

func main() {
	g := game.Create()
	g.Guess("CRANE", [game.WORD_LEN]int{1, 1, 0, 0, 1})
	g.Guess("MERCY", [game.WORD_LEN]int{0, 1, 1, 1, 0})
	g.Guess("ULCER", [game.WORD_LEN]int{2, 2, 2, 2, 2})
	// g.Guess("CRAVE", [game.WORD_LEN]int{0, 2, 2, 0, 2})
	// g.Guess("ERASE", [game.WORD_LEN]int{0, 2, 2, 0, 2})
	// g.Guess("FRAME", [game.WORD_LEN]int{2, 2, 2, 2, 2})
	g.Print()
}
