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
	words   []string
}

func Create() Game {
	return Game{[GUESS_COUNT]guess{}, 0, Words()}
}

func (g *Game) isValidWord(word string) bool {
	for i := 0; i < len(g.words); i++ {
		if strings.EqualFold(strings.ToUpper(word), strings.ToUpper(g.words[i])) {
			return true
		}
	}
	return false
}

func (g *Game) Guess(word string, scores [WORD_LEN]int) error {
	if len(word) != WORD_LEN {
		return errors.New("only 5-character strings are allowed")
	}
	if g.n >= GUESS_COUNT {
		return errors.New("guess limit reached")
	}
	if !g.isValidWord(word) {
		return errors.New("invalid word")
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
