package main

import (
	"fmt"
	"github.com/loganjspears/joker/hand"
	"log"
	"net/http"
)

var BotName = "GOd of Gamblers"
var listenPort = ":5000"

func init() {
	// not interested in timestamps for logging
	log.SetFlags(0)
}

// this is an HTTP bot server for MachinePoker
func main() {
	http.HandleFunc("/bot/gog", botHandler)
	http.ListenAndServe(listenPort, nil)
}

// main handler, triggered by poker server
func botHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		registerBot(w)
	case "POST":
		game := ReadGame(r.Body)
		DisplayGame(game)

		var bet int
		if game.State != "complete" {
			bet = play(game)
		}
		fmt.Fprintf(w, "{\"bet\": \"%d\"}", bet)
	default:
		log.Fatal("Method unsupported:", r.Method)
	}
}

func registerBot(w http.ResponseWriter) {
	fmt.Fprintf(w, "{\"info\": { \"name\": \"%s\" } }", BotName)
}

func play(game *Game) int {
	// consider all cards when calculating odds
	all := append(game.Community, game.Self.Cards...)
	myCards := Cards(all)

	// convert to joker hand and calculate ranking
	myHand := hand.New(myCards)
	logger.Println("** myHand:", myHand)

	// Note: printed value of rank is wrong, subtract 1
	// fmt.Printf("ranking: %s\n", myHand.Ranking()-1)

	switch game.State {
	case "pre-flop":
		return calculatePreflopBet(game, myHand)
	case "flop":
		return calculateBet(game, myHand)
	case "turn":
		return calculateBet(game, myHand)
	case "river":
		return calculateBet(game, myHand)
	default:
		log.Fatal("Undefined game state:", game.State)
		return -1
	}
}

func calculatePreflopBet(game *Game, myHand *hand.Hand) int {
	if myHand.Ranking() == hand.Pair {
		return raise(game)
	}
	return call(game)
}

func calculateBet(game *Game, myHand *hand.Hand) int {
    if safeguard(game, myHand) {
        if myHand.Ranking() >= hand.TwoPair {
            return raise(game)
        } else if (myHand.Ranking() >= hand.Pair || game.Self.Wagered > 50) {
            return call(game)
        }
    }
	return fold(game)
}

func raise(game *Game) int {
	if game.Betting.CanRaise {
		logger.Println("-> raising:", game.Betting.Raise)
		return game.Betting.Raise
	}
	return call(game)
}

func call(game *Game) int {
	logger.Println("-> calling:", game.Betting.Call)
	return game.Betting.Call
}

func fold(game *Game) int {
	logger.Println("-> folding")
	return 0
}

func safeguard(game *Game, myHand *hand.Hand) bool {
	if game.Betting.Call < 100 {
		return true
	}
    if myHand.Ranking() >= hand.ThreeOfAKind {
        return true
    }
    return false
}
