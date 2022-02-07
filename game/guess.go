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

func (g *guess) print() {
	reset := "\x1b[0m"
	colours := [3]string{"\x1b[40m", "\x1b[43m", "\x1b[42m"}
	for i, x := range g.word {
		print(colours[g.score[i]], " ", string(x), " ")
	}
	print(reset, "\n")
}
