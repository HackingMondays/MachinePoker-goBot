package main

import (
	"testing"
	"os"
	"log"
    "fmt"
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

//func TestCard2Joker(t *testing.T) {
//	fixture := [...]string { "2c", "Qh", "Jc" }
//	// expectd := [...]string { "2♣", "Q♥", "J♣" }
//
//	for _, fix := range fixture {
//		fmt.Println(Card2Joker(fix))
//	}
//}

func TestAppend(t *testing.T) {
    s1 := []string{"test", "test1"}
    s2 := []string{"test2", "test3"}
    all := append(s1, s2...)
    if len(all) != 4 {
        t.Error("len(all)=", len(all))
    }
}

func TestCardRanking(t *testing.T) {
    myCards := Cards([] string { "3c", "3h", "2s", "4s", "5c" })
    myHand := hand.New(myCards)
    rank := myHand.Ranking()

    fmt.Println(myHand)

    // The printed rank is wrong, need to subtract 1
    fmt.Println(rank-1)

    if rank == hand.Pair {
        fmt.Println("ok, is a pair\n")
    } else {
        t.Error("Hand ranking expected: 'Pair', but was", rank-1)
    }
}
