package main

import (
	"fmt"
	"io"
	"encoding/json"
	"strings"
	"github.com/loganjspears/joker/hand"
)

type Game struct {
	Community [] string
	State     	 string
	Hand         int
	Betting      Betting
	Self    	 Self
}

type Betting struct {
	Call     int
	Raise    int
	CanRaise bool
}

type Self struct {
	Name	string
	Blind	int
	Ante	int
	Wagered int
	State	string
	Chips 	int
	Actions map[string] [] *Action
	Cards 	[]  string
	Position int
	Brain   [] *string
}

type Action struct {
	Type	string
	Bet 	int
}

func ReadGame(reader io.Reader) *Game {
	var game *Game
	json.NewDecoder(reader).Decode(&game)
	return game
}

func Display (game *Game) {
    fmt.Println("\n--- game -------------")
	fmt.Printf("community: ")
	for _, community := range game.Community {
		fmt.Printf("%s, ", community)
	}
    fmt.Printf("\ncards: ")
    for _, card := range game.Self.Cards {
        fmt.Printf("%s, ", card)
    }
	fmt.Printf("\nstate: %s\nhand: %d", game.State, game.Hand)
	fmt.Printf("\nbetting: %d, %d, %t", game.Betting.Call, game.Betting.Raise, game.Betting.CanRaise)
	fmt.Printf("\nself: %s, %d, %d, %d, %s", game.Self.Name, game.Self.Blind, game.Self.Ante, game.Self.Wagered, game.Self.State)
	fmt.Printf("\nactions: ")
	for _, action := range game.Self.Actions["pre-flop"] {
		fmt.Printf("\ttype: %s, bet: %d", action.Type, action.Bet)
	}
}



// convert JSPoker cards to Joker cards
// https://github.com/loganjspears/joker/blob/master/hand/card.go
func Card2Joker(mycard string) *hand.Card {

	var joker string = mycard
	source := 	   [4]string {"s","h","d","c"}
	destination := [4]string {"♠","♥","♦","♣"}

	for i := 0; i < 4; i++ {
		joker = strings.Replace(joker, source[i], destination[i], 1)
	}
	cr := hand.AceSpades
	cr.UnmarshalText([]byte(joker))
	return cr
}
