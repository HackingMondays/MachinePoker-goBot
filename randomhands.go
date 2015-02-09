package main

import (
	"fmt"
    "log"
	"net/http"
	"github.com/loganjspears/joker/hand"
)

var BotName = "GOd of Gamblers"

func main() {
	http.HandleFunc("/bot/gog", botHandler)
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func botHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            fmt.Fprintf(w, "{\"info\": { \"name\": \"%s\" } }", BotName)
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

func play(game *Game) int {

    // consider all cards when calculating odds
    all := append(game.Community, game.Self.Cards...)
	myCards := Cards(all)

    // convert to joker hand and calculate ranking
    myHand := hand.New(myCards)
    fmt.Println("** myHand:", myHand)

    // TODO: printed value of rank is wrong, subtract 1
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
            return -1;
    }
}

func calculatePreflopBet(game *Game, myHand *hand.Hand) int {
    if myHand.Ranking() == hand.Pair {
        return raiseOrCall(game)
    } else {
        return call(game)
    }
}

func calculateBet(game *Game, myHand *hand.Hand) int {
    if myHand.Ranking() >= hand.TwoPair {
        return raiseOrCall(game)
    } else if myHand.Ranking() >= hand.Pair || game.Self.Wagered > 30 {
        return call(game)
    }
    return fold(game)
}

func raiseOrCall(game *Game) int {
	if game.Betting.CanRaise {
        fmt.Println("-> raising:", game.Betting.Raise)
		return game.Betting.Raise
	} else {
		return call(game)
	}
}

func call (game *Game) int {
    fmt.Println("-> calling:", game.Betting.Call)
    return game.Betting.Call
}

func fold (game *Game) int {
    fmt.Println("-> folding")
    return 0
}