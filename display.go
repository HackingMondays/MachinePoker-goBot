package main

// Utility function to dump a game object to the console
// Requires an instance of *log.Logger (cf. logger.go)
func DisplayGame(game *Game) {
	logger.Println("\n--- game -------------")
	DisplayCards("community", game.Community)
	DisplayCards("cards", game.Self.Cards)
	logger.Printf("state: %s\nhand: %d\n", game.State, game.Hand)
	logger.Printf("betting: call=%d, raise=%d, canRaise=%t\n", game.Betting.Call, game.Betting.Raise, game.Betting.CanRaise)
	DisplaySelf(&game.Self)
	DisplayPlayers(game.Players)
}

func DisplayCards(label string, cards []string) {
	logger.Printf("%s: ", label)
	for _, card := range cards {
		logger.Printf("%s, ", card)
	}
	logger.Println()
}

func DisplaySelf(self *Self) {
	logger.Println("self: ")
	logger.Printf("  name: %s, blind: %d, ante: %d, wagered: %d, state: %s, chips: %d\n",
		self.Name, self.Blind, self.Ante, self.Wagered, self.State, self.Chips)
	DisplayActions(self.Actions)
}

func DisplayActions(actionMap map[string][]*Action) {
	logger.Println("  actions:")
	for key, actions := range actionMap {
		logger.Printf("    state: %s\n", key)
		for _, action := range actions {
			logger.Printf("      type: %s, bet: %d\n", action.Type, action.Bet)
		}
	}
}

func DisplayPlayers(players []Player) {
	logger.Println("players:")
	for _, player := range players {
		if player.Name != BotName {
			DisplayPlayer(player)
		}
	}
	logger.Println()
}

func DisplayPlayer(player Player) {
	logger.Printf("  name: %s, blind: %d, ante: %d, wagered: %d, state: %s, chips: %d\n",
		player.Name, player.Blind, player.Ante, player.Wagered, player.State, player.Chips)
	DisplayActions(player.Actions)
}
