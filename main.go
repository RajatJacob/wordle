package main

import (
	"fmt"

	"github.com/RajatJacob/wordle/game"
)

func main() {
	fmt.Println("wordle-bot")
	g := game.Create()
	g.Guess("WORDS")
	g.Print()
}
