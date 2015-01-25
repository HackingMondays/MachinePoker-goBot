package main

import (
	"testing"
	"github.com/loganjspears/joker/hand"
	"github.com/loganjspears/joker/jokertest"
)

// copied from hand_test.go
// but removed the . in front of import
// and prefixed Card and New with hand.
//
// see: http://stackoverflow.com/questions/6478962/what-does-the-dot-or-period-in-a-go-import-statement-do

type equality int

const (
	greaterThan equality = iota
	lessThan
	equalTo
)

type testEquality struct {
	cards1 []*hand.Card
	cards2 []*hand.Card
	e      equality
}

var equalityTests = []testEquality{
	{
		jokertest.Cards("As", "5s", "4s", "3s", "2s"),
		jokertest.Cards("Ks", "Kc", "Kh", "Jd", "Js"),
		greaterThan,
	},
	{
		jokertest.Cards("Ts", "9h", "8d", "7c", "6s", "2h", "3s"),
		jokertest.Cards("Ts", "9h", "8d", "7c", "6s", "Ah", "Ks"),
		equalTo,
	},
}

func TestCompareHands(t *testing.T) {
	for _, test := range equalityTests {
		h1 := hand.New(test.cards1)
		h2 := hand.New(test.cards2)
		compareTo := h1.CompareTo(h2)

		switch test.e {
		case greaterThan:
			if compareTo <= 0 {
				t.Errorf("expected %v to be greater than %v", h1, h2)
			}
		case lessThan:
			if compareTo >= 0 {
				t.Errorf("expected %v to be less than %v", h1, h2)
			}
		case equalTo:
			if compareTo != 0 {
				t.Errorf("expected %v to be equal to %v", h1, h2)
			}
		}
	}
}

