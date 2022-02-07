package main

import (
	"fmt"

	"github.com/RajatJacob/wordle/game"
)

func main() {
	fmt.Println("wordle-bot")
	g := game.Create()
	g.Guess("CRANE", [game.WORD_LEN]int{0, 1, 0, 0, 1})
	g.Guess("OTHER", [game.WORD_LEN]int{0, 0, 0, 2, 2})
	g.Guess("SUPER", [game.WORD_LEN]int{0, 0, 0, 2, 2})
	g.Guess("WIDER", [game.WORD_LEN]int{0, 0, 2, 2, 2})
	g.Guess("MEMES", [game.WORD_LEN]int{0, 1, 0, 0, 2})
	g.Guess("ELDER", [game.WORD_LEN]int{2, 2, 2, 2, 2})
	g.Print()
}
