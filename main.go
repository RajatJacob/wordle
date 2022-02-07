package main

import (
	"fmt"

	"github.com/RajatJacob/wordle/game"
)

func main() {
	fmt.Println("wordle-bot")
	g := game.Create()
	g.Print()
	g.Guess("WORDS", [game.WORD_LEN]int{0, 1, 2, 1, 2})
	g.Print()
}
