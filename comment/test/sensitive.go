package main

import (
	"comment/sensitiveWord"
	"log"
)

func main() {
	word := "你是sb"
	sensitiveWord.InitWords()
	Word := sensitiveWord.ToInsensitive(word)
	log.Fatal(Word)
}
