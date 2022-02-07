package game

import (
	"errors"
	"strings"
)

const WORD_LEN int = 5
const GUESS_COUNT int = 6

type Game struct {
	guesses [GUESS_COUNT]string
	n       int
}

func Create() Game {
	return Game{[GUESS_COUNT]string{}, 0}
}

func (g *Game) Guess(guess string) error {
	if len(guess) != WORD_LEN {
		return errors.New("only 5-character strings are allowed")
	}
	if g.n >= GUESS_COUNT {
		return errors.New("guess limit reached")
	}
	g.guesses[g.n] = guess
	g.n++
	return nil
}

func (g *Game) Print() {
	print(strings.Join(g.guesses[:], ", "), "\n", g.n, "\n\n")
}
