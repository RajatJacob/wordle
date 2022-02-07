package game

import (
	"errors"
	"strings"
)

const WORD_LEN int = 5
const GUESS_COUNT int = 6

type Game struct {
	guesses [GUESS_COUNT]guess
	n       int
}

func Create() Game {
	return Game{[GUESS_COUNT]guess{}, 0}
}

func (g *Game) Guess(word string, scores [WORD_LEN]int) error {
	if len(word) != WORD_LEN {
		return errors.New("only 5-character strings are allowed")
	}
	if g.n >= GUESS_COUNT {
		return errors.New("guess limit reached")
	}
	g.guesses[g.n] = guess{strings.ToUpper(word), scores}
	g.n++
	return nil
}

func (g *Game) Print() {
	for i := 0; i < g.n; i++ {
		g.guesses[i].print()
	}
}
