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
            DisplayGame(game)

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

    // TODO: printed value of rank is wrong, subract 1
    fmt.Printf("ranking: %s\n", myHand.Ranking()-1)

    // strategy
	if game.State == "pre-flop" {
		if myHand.Ranking() == hand.Pair {
			ret = raise(game)
		} else {
			ret = rand.Intn(2) * game.Betting.Call
		}
	} else {
        if myHand.Ranking() >= hand.ThreeOfAKind {
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
