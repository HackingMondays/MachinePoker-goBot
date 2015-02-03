package main

import (
	"fmt"
    "log"
	"net/http"
	"math/rand"
	"github.com/loganjspears/joker/hand"
)

func main() {
	http.HandleFunc("/bot/gog", botHandler)
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func botHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            fmt.Fprintf(w, "{\"info\": { \"name\": \"GOd of Gamblers\" } }")
        case "POST":
            game := ReadGame(r.Body)
            Display(game)

            var bet int
            if game.State != "complete" {
                bet = play(game)
            }
            fmt.Fprintf(w, "%d", bet)
        default:
            log.Fatal("Method unsupported")
    }
}

func play(game *Game) int {
	var ret int

    // consider all cards to calculate odds
    all := append(game.Community, game.Self.Cards...)
	myCards := Cards(all)

    // convert to joker hand and calculate ranking
    myHand := hand.New(myCards)
    fmt.Println("\n** Hand: ")
    fmt.Println(myHand)

    // TODO: ranking is currently wrong, needs calc ?
    fmt.Printf("ranking: %s\n", myHand.Ranking())

    // strategy
	if game.State == "pre-flop" {
		if myHand.Ranking() == hand.Pair {
			ret = raise(game)
		} else {
			ret = rand.Intn(2) * game.Betting.Call
		}
	} else {
        if myHand.Ranking() >= hand.Flush {
            ret = raise(game)
        } else if myHand.Ranking() == hand.Pair {
            ret = game.Betting.Call
        }
	}
	return ret
}

func raise(game *Game) int {
    fmt.Println("raising")
	if game.Betting.CanRaise {
		return game.Betting.Raise
	} else {
		return game.Betting.Call
	}
}


// ----------------------------------------

// copied from jokertest.go
func Cards(list []string) []*hand.Card {
	cards := []*hand.Card{}
	for _, s := range list {
		cards = append(cards, card(s))
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
	rankMap = map[string] hand.Rank {
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

	suitMap = map[string] hand.Suit {
	"s": hand.Spades,
	"h": hand.Hearts,
	"d": hand.Diamonds,
	"c": hand.Clubs,
    }
)


