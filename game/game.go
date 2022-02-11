package game

import (
	"errors"
	"strings"
)

const WORD_LEN int = 5
const GUESS_COUNT int = 6

type Game struct {
	guesses []guess
	words   set
	letters map[rune]int
}

func Create() Game {
	return Game{[]guess{}, Words(), make(map[rune]int)}
}

func (g *Game) Guess(word string, scores [WORD_LEN]int) error {
	if len(word) != WORD_LEN {
		return errors.New("only 5-character strings are allowed")
	}
	if len(g.guesses) >= GUESS_COUNT {
		return errors.New("guess limit reached")
	}
	if !g.words.has(word) {
		return errors.New("invalid word")
	}
	g.guesses = append(g.guesses, guess{strings.ToUpper(word), scores})
	for i := range word {
		colour, ok := g.letters[rune(word[i])]
		if !ok || colour < scores[i] {
			g.letters[rune(word[i])] = scores[i]
		}
	}
	return nil
}

func (g *Game) Print() {
	for _, guess := range g.guesses {
		guess.print()
	}
	words := g.possibleWords().elements()
	println("\n")
	println(len(words), "words\n")
	if len(words) > 20 {
		words = words[:20]
	}
	println(strings.Join(words, "\n"))
	println(len(g.letters), "/", 26, "\n")
	for l, c := range g.letters {
		printColouredLetter(l, c)
	}
	println()
}

func includes(word string, char byte) bool {
	for _, s := range word {
		if s == rune(char) {
			return true
		}
	}
	return false
}

func (g *Game) removeRemoveInvalidWords(guess *guess, i int) {
	for _, word := range g.words.elements() {
		// colour, ok := g.letters[rune(word[i])]
		cond := map[int]bool{
			BLACK:  !includes(word, guess.word[i]),
			YELLOW: word[i] != guess.word[i] && includes(word, guess.word[i]),
			GREEN:  word[i] == guess.word[i],
		}[guess.score[i]]
		if !cond {
			g.words.delete(word)
		}
	}
}

func (g *Game) possibleWords() *set {
	for _, guess := range g.guesses {
		for i := range guess.score {
			g.removeRemoveInvalidWords(&guess, i)
		}
	}
	return &g.words
}
