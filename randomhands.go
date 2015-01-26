package main

import (
	"fmt"
	"net/http"
	"github.com/loganjspears/joker/hand"
)
// https://golang.org/doc/articles/wiki/
func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/play/", playHandler)
	http.ListenAndServe(":8080", nil)
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

