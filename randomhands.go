package main

import (
	"fmt"
	"github.com/loganjspears/joker/hand"
)

func main() {
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

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
