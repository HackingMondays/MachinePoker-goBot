package main

import (
    "flag"
	"fmt"
	"log"
	"net/http"
)

var botName = "GOd of Gamblers"
var listenPort = ":5000"
var pokerPlayer defaultPlayer

// define command line parameters
func init() {
    flag.StringVar(&botName, "name", botName, "name of the bot")
    flag.StringVar(&listenPort, "port", listenPort, "listen port, eg. ':5000'")
}

// this is an HTTP bot server for MachinePoker
func main() {
    // parse command line flags
    flag.Parse()

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
        fmt.Fprintf(w, "{\"bet\": \"%d\"}", betForGame(game))
	default:
		log.Fatal("Method unsupported:", r.Method)
	}
}

// registers bot with server
func registerBot(w http.ResponseWriter) {
	fmt.Fprintf(w, "{\"info\": { \"name\": \"%s\" } }", botName)
}

// return bet to server of display completed game
func betForGame(game *Game) int {
    if game.State != "complete" {
        return pokerPlayer.Play(game)
    }
    logger.Println(game)
    return 0;
}
