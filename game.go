package main

import (
	"encoding/json"
	"io"
)

// The game object, reverse-engineered from PokerJS JSON
type Game struct {
	Community []string
	State     string
	Hand      int
	Betting   Betting
	Self      Self
	Players   []Player
}

type Betting struct {
	Call     int
	Raise    int
	CanRaise bool
}

type Self struct {
	Name     string
	Blind    int
	Ante     int
	Wagered  int
	State    string
	Chips    int
	Actions  map[string][]*Action
	Cards    []string
	Position int
	Brain    []string
}

type Action struct {
	Type string
	Bet  int
}

type Player struct {
	Name    string
	Blind   int
	Ante    int
	Wagered int
	State   string
	Chips   int
	Actions map[string][]*Action
}

func ReadGame(reader io.Reader) *Game {
	var game *Game
	json.NewDecoder(reader).Decode(&game)
	return game
}
