package main

import (
	"github.com/RajatJacob/wordle/game"
)

func main() {
	g := game.Create()
	g.Guess("WATER", [game.WORD_LEN]int{0, 1, 0, 1, 1})
	g.Guess("LEARN", [game.WORD_LEN]int{0, 1, 2, 1, 0})
	g.Guess("GRAPE", [game.WORD_LEN]int{0, 2, 2, 0, 2})
	g.Guess("CRAVE", [game.WORD_LEN]int{0, 2, 2, 0, 2})
	g.Guess("ERASE", [game.WORD_LEN]int{0, 2, 2, 0, 2})
	g.Guess("FRAME", [game.WORD_LEN]int{2, 2, 2, 2, 2})
	g.Print()
}
