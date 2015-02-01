package main

import (
	"testing"
	"os"
	"log"
)

const dataFile = "data/data.json"

// read JSON from file
func readGameFromFile(fileName string) Game {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return readGame(file)
}

func TestReadGame(t *testing.T) {
	var game Game
	game = readGameFromFile(dataFile)
	if game.Community == nil {
		t.Errorf("cannot read game")
	}
	Display(&game)
}

