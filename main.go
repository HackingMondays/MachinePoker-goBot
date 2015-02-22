package main

import (
	"fmt"
	"log"
	"net/http"
)

var botName = "GOd of Gamblers"
var listenPort = ":5000"
var pokerPlayer defaultPlayer

// this is an HTTP bot server for MachinePoker
func main() {
    // set default logger
    logger = Info

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
			bet = pokerPlayer.Play(game)
		}
		fmt.Fprintf(w, "{\"bet\": \"%d\"}", bet)
	default:
		log.Fatal("Method unsupported:", r.Method)
	}
}

func registerBot(w http.ResponseWriter) {
	fmt.Fprintf(w, "{\"info\": { \"name\": \"%s\" } }", botName)
}
