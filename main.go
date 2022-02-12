package main

import (
	"fmt"

	"github.com/RajatJacob/wordle/game"
)

func main() {
	g := game.Create()
	var e error = nil
	for e == nil {
		var word string
		var scores [game.WORD_LEN]int
		fmt.Scanf("%s", &word)
		for i := range scores {
			fmt.Scanf("%d", &scores[i])
		}
		e = g.Guess(word, scores)
		if e != nil {
			panic(e)
		}
		g.Print()
	}
}
