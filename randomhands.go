package main

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"github.com/loganjspears/joker/hand"
	"math/rand"
)


func main() {
	http.HandleFunc("/", handler)
	// http.ListenAndServe(":8081", nil)
	http.ListenAndServe("0.0.0.0:8081", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
	var game Game
	var ret int

	fmt.Printf("Method: %s\n", r.Method);

	json.NewDecoder(r.Body).Decode(&game)
	Display(&game)

	if r.Method == "GET" {
		fmt.Fprintf(w, "{\"info\": { \"name\": \"GOd of Gamblers\" } }")
	} else {
		if game.Betting.CanRaise {
			ret = rand.Intn(2) * game.Betting.Raise
		} else {
			ret = game.Betting.Call
		}
		fmt.Fprintf(w, "%d", ret)
	}
}

// example read JSON from file
func decodeFromFile() {
	var game Game
	file, err := os.Open(dataFile)
	if err != nil {
		// return nil, err
		fmt.Println("Error:", err)
	}
	err = json.NewDecoder(file).Decode(&game)

	defer file.Close()

	Display(&game)
}

// https://golang.org/doc/articles/wiki/
func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// same as Play() but write to ResponseWriter
func playHandler(w http.ResponseWriter, r *http.Request) {
	deck := hand.NewDealer().Deck()
	h1 := hand.New(deck.PopMulti(5))
	h2 := hand.New(deck.PopMulti(5))
	winner := FindWinner(h1, h2)
	fmt.Fprintf(w,"Winner is: %s", winner)
}

// not used
func Play() {
	deck := hand.NewDealer().Deck()
	h1 := hand.New(deck.PopMulti(5))
	h2 := hand.New(deck.PopMulti(5))
	winner := FindWinner(h1, h2)
	fmt.Println("Winner is:", winner)
}

func FindWinner(h1 *hand.Hand, h2 *hand.Hand) []*hand.Card {
	fmt.Println(h1)
	fmt.Println(h2)
	hands := hand.Sort(hand.SortingHigh, hand.DESC, h1, h2)
	return hands[0].Cards()
}

