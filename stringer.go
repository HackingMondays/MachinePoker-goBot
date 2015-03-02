package main

import (
	"bytes"
	"fmt"
	"strings"
)

func (game *Game) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[--- game ---]")
	buffer.WriteString("\ncards (self/community): ")
	buffer.WriteString(game.Self.Cards.String())
	buffer.WriteString(" / ")
	buffer.WriteString(game.Community.String())
	buffer.WriteString(fmt.Sprintf("\nstate: %s, hand: %d, ", game.State, game.Hand))
	buffer.WriteString(game.Betting.String())
	buffer.WriteString("players:\n")
	for _, player := range game.Players {
		buffer.WriteString(player.String())
	}
	return buffer.String()
}

func (cards GameCards) String() string {
	return strings.Join(cards, ", ")
}

func (betting *Betting) String() string {
	return fmt.Sprintf("call=%d, raise=%d, canRaise=%t\n",
		betting.Call, betting.Raise, betting.CanRaise)
}

func (action *Action) String() string {
	return fmt.Sprintf("      type: %s, bet: %d\n",
		action.Type, action.Bet)
}

func (player *Player) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("  name: %s,", player.Name))
	buffer.WriteString(fmt.Sprintf(" blind: %d,", player.Blind))
	buffer.WriteString(fmt.Sprintf(" ante: %d,", player.Ante))
	buffer.WriteString(fmt.Sprintf(" wagered: %d,", player.Wagered))
	buffer.WriteString(fmt.Sprintf(" state: %s,", player.State))
	buffer.WriteString(fmt.Sprintf(" chips: %d\n", player.Chips))

	for key, actions := range player.Actions {
		buffer.WriteString(fmt.Sprintf("    state: %s\n", key))
		for _, action := range actions {
			buffer.WriteString(fmt.Sprintf("      type: %s,", action.Type))
			buffer.WriteString(fmt.Sprintf(" bet: %d\n", action.Bet))
		}
	}
	return buffer.String()
}
