package main

import (
	"fmt"
	"io"
	"encoding/json"
)

type Game struct {
	Community [] string
	State     	 string
	Hand         int
	Betting      Betting
	Self    	 Self
    Players []   Player
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
	Brain   [] string
}

type Action struct {
	Type	string
	Bet 	int
}

type Player struct {
    Name	string
    Blind	int
    Ante	int
    Wagered int
    State	string
    Chips 	int
    Actions map[string] [] *Action
}

func ReadGame(reader io.Reader) *Game {
	var game *Game
	json.NewDecoder(reader).Decode(&game)
	return game
}

func DisplayGame(game *Game) {
    fmt.Println("\n--- game -------------")
    DisplayCards("community", game.Community)
    DisplayCards("cards", game.Self.Cards)
	fmt.Printf("state: %s\nhand: %d\n", game.State, game.Hand)
	fmt.Printf("betting: call=%d, raise=%d, canRaise=%t\n", game.Betting.Call, game.Betting.Raise, game.Betting.CanRaise)
    DisplaySelf(&game.Self)
    DisplayPlayers(game.Players)
}

func DisplayCards(label string, cards []string) {
    fmt.Printf("%s: ", label)
    for _, card := range cards {
        fmt.Printf("%s, ", card)
    }
    fmt.Println()
}

func DisplaySelf(self *Self) {
    fmt.Println("self: ")
    fmt.Printf("  name: %s, blind: %d, ante: %d, wagered: %d, state: %s, chips: %d\n",
        self.Name, self.Blind, self.Ante, self.Wagered, self.State, self.Chips)
    DisplayActions(self.Actions)
}

func DisplayActions(actionMap map[string] [] *Action) {
    fmt.Println("  actions:")
    for key, actions := range actionMap {
        fmt.Printf("    state: %s\n", key)
        for _, action := range actions {
            fmt.Printf("      type: %s, bet: %d\n", action.Type, action.Bet)
        }
    }
}

func DisplayPlayers(players [] Player) {
    fmt.Println("players:")
    for _, player := range players {
        if player.Name != BotName {
            DisplayPlayer(player)
        }
    }
    fmt.Println()
}

func DisplayPlayer(player Player) {
    fmt.Printf("  name: %s, blind: %d, ante: %d, wagered: %d, state: %s, chips: %d\n",
    player.Name, player.Blind, player.Ante, player.Wagered, player.State, player.Chips)
    DisplayActions(player.Actions)
}