package main

import (
	"fmt"
	"io"
	"encoding/json"
	// "github.com/loganjspears/joker/hand"
)

type Game struct {
	Community [] *string
	State     	  string
	Hand          int
	Betting       Betting
	Self		  Self
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
	Actions [] *Action
	Cards 	[] *string
	Position int
	Brain   [] *string
}

type Action struct {
	// "actions": { "pre-flop": [ { "type": "call", "bet": 5 } ] },
}

func readGame(reader io.Reader) Game {
	var game Game
	json.NewDecoder(reader).Decode(&game)
	return game
}

func Display (game *Game) {
	fmt.Printf("community: ")
	for _,card := range game.Community {
		fmt.Printf("%s,", *card)
	}
	fmt.Printf("\nstate: %s\nhand: %d", game.State, game.Hand)
	fmt.Printf("\nbetting: %d, %d, %t", game.Betting.Call, game.Betting.Raise, game.Betting.CanRaise)
	fmt.Printf("\nself: %s, %d, %d, %d, %s", game.Self.Name, game.Self.Blind, game.Self.Ante, game.Self.Wagered, game.Self.State)
}

