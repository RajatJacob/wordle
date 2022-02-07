package game

import (
	"log"
	"os"
	"strings"
)

func Words() set {
	content, err := os.ReadFile("./game/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(strings.ToUpper(string(content[:])), "\n")
	for i := range words {
		words[i] = strings.ToUpper(words[i])
	}
	set := newSet(words)
	return set
}
