package game

import (
	"log"
	"os"
	"strings"
)

func Words() []string {
	content, err := os.ReadFile("./game/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(content[:]), "\n")
	return words
}
