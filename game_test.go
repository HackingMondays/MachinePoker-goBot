package main

import (
	"testing"
	"os"
	"log"
)

const dataFile = "data/data.json"

// read JSON from file
func ReadGameFromFile(fileName string) *Game {
	file, err := os.Open(fileName)
	if err != nil {
		// this stops all further processing
		log.Fatal(err)
	}
	defer file.Close()
	return ReadGame(file)
}

func TestReadGame(t *testing.T) {
	var game *Game
	game = ReadGameFromFile(dataFile)
	Display(game)
}

