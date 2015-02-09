package main

import (
	"log"
)

func DisplayGame(game *Game) {
	log.Println("\n--- game -------------")
	DisplayCards("community", game.Community)
	DisplayCards("cards", game.Self.Cards)
	log.Printf("state: %s\nhand: %d\n", game.State, game.Hand)
	log.Printf("betting: call=%d, raise=%d, canRaise=%t\n", game.Betting.Call, game.Betting.Raise, game.Betting.CanRaise)
	DisplaySelf(&game.Self)
	DisplayPlayers(game.Players)
}

func DisplayCards(label string, cards []string) {
	log.Printf("%s: ", label)
	for _, card := range cards {
		log.Printf("%s, ", card)
	}
	log.Println()
}

func DisplaySelf(self *Self) {
	log.Println("self: ")
	log.Printf("  name: %s, blind: %d, ante: %d, wagered: %d, state: %s, chips: %d\n",
		self.Name, self.Blind, self.Ante, self.Wagered, self.State, self.Chips)
	DisplayActions(self.Actions)
}

func DisplayActions(actionMap map[string][]*Action) {
	log.Println("  actions:")
	for key, actions := range actionMap {
		log.Printf("    state: %s\n", key)
		for _, action := range actions {
			log.Printf("      type: %s, bet: %d\n", action.Type, action.Bet)
		}
	}
}

func DisplayPlayers(players []Player) {
	log.Println("players:")
	for _, player := range players {
		if player.Name != BotName {
			DisplayPlayer(player)
		}
	}
	log.Println()
}

func DisplayPlayer(player Player) {
	log.Printf("  name: %s, blind: %d, ante: %d, wagered: %d, state: %s, chips: %d\n",
		player.Name, player.Blind, player.Ante, player.Wagered, player.State, player.Chips)
	DisplayActions(player.Actions)
}
