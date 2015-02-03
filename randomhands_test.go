package main

import (
	"encoding/json"
	"testing"

	"github.com/loganjspears/joker/hand"
	"github.com/loganjspears/joker/jokertest"
)

type testPair struct {
	cards       []*hand.Card
	arrangement []*hand.Card
	ranking     hand.Ranking
	description string
}

var tests = []testPair{
	{
		jokertest.Cards("Ks", "Qs", "Js", "As", "9d"),
		jokertest.Cards("As", "Ks", "Qs", "Js", "9d"),
		hand.HighCard,
		"high card ace high",
	},
	{
		jokertest.Cards("Ks", "Qh", "Qs", "Js", "9d"),
		jokertest.Cards("Qh", "Qs", "Ks", "Js", "9d"),
        hand.Pair,
		"pair of queens",
	},
	{
		jokertest.Cards("2s", "Qh", "Qs", "Js", "2d"),
		jokertest.Cards("Qh", "Qs", "2s", "2d", "Js"),
        hand.TwoPair,
		"two pair queens and twos",
	},
	{
		jokertest.Cards("6s", "Qh", "Ks", "6h", "6d"),
		jokertest.Cards("6s", "6h", "6d", "Ks", "Qh"),
        hand.ThreeOfAKind,
		"three of a kind sixes",
	},
	{
		jokertest.Cards("Ks", "Qs", "Js", "As", "Td"),
		jokertest.Cards("As", "Ks", "Qs", "Js", "Td"),
        hand.Straight,
		"straight ace high",
	},
	{
		jokertest.Cards("2s", "3s", "4s", "As", "5d"),
		jokertest.Cards("5d", "4s", "3s", "2s", "As"),
        hand.Straight,
		"straight five high",
	},
	{
		jokertest.Cards("7s", "4s", "5s", "3s", "2s"),
		jokertest.Cards("7s", "5s", "4s", "3s", "2s"),
        hand.Flush,
		"flush seven high",
	},
	{
		jokertest.Cards("7s", "7d", "3s", "3d", "7h"),
		jokertest.Cards("7s", "7d", "7h", "3s", "3d"),
        hand.FullHouse,
		"full house sevens full of threes",
	},
	{
		jokertest.Cards("7s", "7d", "3s", "7c", "7h"),
		jokertest.Cards("7s", "7d", "7c", "7h", "3s"),
        hand.FourOfAKind,
		"four of a kind sevens",
	},
	{
		jokertest.Cards("Ks", "Qs", "Js", "Ts", "9s"),
		jokertest.Cards("Ks", "Qs", "Js", "Ts", "9s"),
        hand.StraightFlush,
		"straight flush king high",
	},
	{
		jokertest.Cards("As", "5s", "4s", "3s", "2s"),
		jokertest.Cards("5s", "4s", "3s", "2s", "As"),
        hand.StraightFlush,
		"straight flush five high",
	},
	{
		jokertest.Cards("As", "Ks", "Qs", "Js", "Ts"),
		jokertest.Cards("As", "Ks", "Qs", "Js", "Ts"),
        hand.RoyalFlush,
		"royal flush",
	},
	{
		jokertest.Cards("As", "Ks", "Qs", "2s", "2c", "2h", "2d"),
		jokertest.Cards("2s", "2c", "2h", "2d", "As"),
        hand.FourOfAKind,
		"four of a kind twos",
	},
}

func TestHands(t *testing.T) {
	for _, test := range tests {
		h := hand.New(test.cards)
		if h.Ranking() != test.ranking {
			t.Fatalf("expected %v got %v", test.ranking, h.Ranking())
		}
		for i := 0; i < 5; i++ {
			actual, expected := h.Cards()[i], test.arrangement[i]
			if actual.Rank() != expected.Rank() || actual.Suit() != expected.Suit() {
				t.Fatalf("expected %v got %v", expected, actual)
			}
		}
		if test.description != h.Description() {
			t.Fatalf("expected \"%v\" got \"%v\"", test.description, h.Description())
		}
	}
}

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

type testOptionsPairs struct {
	cards       []*hand.Card
	arrangement []*hand.Card
	options     []func(*hand.Config)
	ranking     hand.Ranking
	description string
}

var optTests = []testOptionsPairs{
	{
		jokertest.Cards("Ks", "Qs", "Js", "As", "9s"),
		jokertest.Cards("As", "Ks", "Qs", "Js", "9s"),
		[]func(*hand.Config){hand.Low},
        hand.Flush,
		"flush ace high",
	},
	{
		jokertest.Cards("7h", "6h", "5s", "4s", "2s", "3s"),
		jokertest.Cards("6h", "5s", "4s", "3s", "2s"),
		[]func(*hand.Config){hand.AceToFiveLow},
        hand.HighCard,
		"high card six high",
	},
	{
		jokertest.Cards("Ah", "6h", "5s", "4s", "2s", "Ks"),
		jokertest.Cards("6h", "5s", "4s", "2s", "Ah"),
		[]func(*hand.Config){hand.AceToFiveLow},
        hand.HighCard,
		"high card six high",
	},
}

func TestHandsWithOptions(t *testing.T) {
	for _, test := range optTests {
		h := hand.New(test.cards, test.options...)
		if h.Ranking() != test.ranking {
			t.Fatalf("expected %v got %v", test.ranking, h.Ranking())
		}
		for i := 0; i < 5; i++ {
			actual, expected := h.Cards()[i], test.arrangement[i]
			if actual.Rank() != expected.Rank() || actual.Suit() != expected.Suit() {
				t.Fatalf("expected %v got %v", expected, actual)
			}
		}
		if test.description != h.Description() {
			t.Fatalf("expected \"%v\" got \"%v\"", test.description, h.Description())
		}
	}
}

//func TestBlanks(t *testing.T) {
//	cards := []*hand.Card{hand.AceSpades}
//	hand := hand.New(cards)
//	if hand.Ranking() != hand.HighCard {
//		t.Fatal("blank card error")
//	}
//
//	cards = []*hand.Card{hand.FiveSpades, hand.FiveClubs}
//	hand = hand.New(cards)
//	if hand.Ranking() != hand.Pair {
//		t.Fatal("blank card error")
//	}
//}

func TestDeck(t *testing.T) {
	deck := hand.NewDealer().Deck()
	if deck.Pop() == deck.Pop() {
		t.Fatal("Two Pop() calls should never return the same result")
	}
	l := len(deck.Cards)
	if l != 50 {
		t.Fatalf("After Pop() deck len = %d; want %d", l, 50)
	}
}

func TestCardJSON(t *testing.T) {
	card := hand.AceSpades

	// to json
	b, err := json.Marshal(card)
	if err != nil {
		t.Fatal(err)
	}

	// and back
	cardCopy := hand.KingHearts
	if err := json.Unmarshal(b, cardCopy); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkHandCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cards := hand.NewDealer().Deck().PopMulti(7)
        hand.New(cards)
	}
}
