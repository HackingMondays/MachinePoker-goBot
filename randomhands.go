package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"github.com/loganjspears/joker/hand"
)

func main() {
	http.HandleFunc("/", nameHandler)
	http.HandleFunc("/bot", botHandler)
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"info\": { \"name\": \"GOd of Gamblers\" } }")
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	game := ReadGame(r.Body)
	Display(game)

	if game.State != "complete" {
		fmt.Fprintf(w, "%d", play(game))
	}
}

func play(game *Game) int {
	var ret int
	myCards := Cards(game.Self.Cards)

	if game.State == "pre-flop" {
		myHand := hand.New(myCards)

		// bet on first hand
		if myHand.Ranking() >= hand.Pair {
			ret = raise(game)
		} else {
			ret = rand.Intn(2) * game.Betting.Call
		}
	} else {
		// in flop, append community cards
		for _, s := range game.Community {
			myCards = append(myCards, card(*s))
		}
		myHand := hand.New(myCards)

		// bet on new hand
		if myHand.Ranking() >= hand.Flush {
			ret = raise(game)
		} else {
			ret = rand.Intn(2) * game.Betting.Call
		}
		ret = raise(game)
	}
	return ret
}

func raise(game *Game) int {
	if game.Betting.CanRaise {
		return game.Betting.Raise
	} else {
		return game.Betting.Call
	}
}


// ----------------------------------------

// copied from jokertest.go
func Cards(list []*string) []*hand.Card {
	cards := []*hand.Card{}
	for _, s := range list {
		cards = append(cards, card(*s))
	}
	return cards
}

func card(s string) *hand.Card {
	if len(s) != 2 {
		panic("jokertest: card string must be two characters")
	}

	rank, ok := rankMap[s[:1]]
	if !ok {
		panic("jokertest: rank not found")
	}

	suit, ok := suitMap[s[1:]]
	if !ok {
		panic("jokertest: suit not found")
	}

	for _, c := range hand.Cards() {
		if rank == c.Rank() && suit == c.Suit() {
			return c
		}
	}
	panic("jokertest: card not found")
}

var (
	rankMap = map[string]hand.Rank{
	"A": hand.Ace,
	"K": hand.King,
	"Q": hand.Queen,
	"J": hand.Jack,
	"T": hand.Ten,
	"9": hand.Nine,
	"8": hand.Eight,
	"7": hand.Seven,
	"6": hand.Six,
	"5": hand.Five,
	"4": hand.Four,
	"3": hand.Three,
	"2": hand.Two,
}

	suitMap = map[string]hand.Suit{
	"s": hand.Spades,
	"h": hand.Hearts,
	"d": hand.Diamonds,
	"c": hand.Clubs,
}
)


