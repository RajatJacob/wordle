package game

const (
	BLACK  = 0
	YELLOW = 1
	GREEN  = 2
)

type guess struct {
	word  string
	score [WORD_LEN]int
}

func printColouredLetter(letter rune, colour int) {
	reset := "\x1b[0m"
	colours := [3]string{"\x1b[40m", "\x1b[43m", "\x1b[42m"}
	print(colours[colour], " ", string(letter), " ", reset)
}

func (g *guess) print() {
	for i, x := range g.word {
		printColouredLetter(x, g.score[i])
	}
	println()
}
