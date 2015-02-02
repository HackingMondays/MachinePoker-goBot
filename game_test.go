package main

import (
	"testing"
	"os"
	"log"
	"fmt"
	"github.com/loganjspears/joker/jokertest"
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

//func TestCard2Joker(t *testing.T) {
//	fixture := [...]string { "2c", "Qh", "Jc" }
//	expectd := [...]string { "2♣", "Q♥", "J♣" }
//
//	for _, fix := range fixture {
//		fmt.Println(Card2Joker(fix))
//	}
//}

func TestCard2Joker(t *testing.T) {
	h1 := jokertest.Cards("2c", "Qh", "Jc")
	fmt.Println(h1)
}