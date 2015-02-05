package main

import (
	"testing"
	"os"
	"log"
    "github.com/loganjspears/joker/hand"
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
    if game.State != "flop" {
        t.Error("Game stated expected: 'flop', but was", game.State)
    }
}

func TestCardRanking(t *testing.T) {
    myCards := Cards([] string { "3c", "3h", "2s" })
    myHand := hand.New(myCards)
    rank := myHand.Ranking()

    // The printed rank is wrong, need to subtract 1
    if rank != hand.Pair {
        t.Error("Hand ranking expected: 'Pair', but was", rank-1)
    }
}

func TestDisplay(t *testing.T) {
    DisplayGame(ReadGameFromFile(dataFile))
}