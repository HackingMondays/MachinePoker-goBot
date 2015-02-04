package main

import (
    "github.com/loganjspears/joker/hand"
)

// converts js-poker cards to joker cards
// (mostly copied from jokertest.go)

func Cards(list []string) []*hand.Card {
    cards := []*hand.Card{}
    for _, s := range list {
        cards = append(cards, card(s))
    }
    return cards
}

func card(s string) *hand.Card {
    if len(s) != 2 {
        panic("jokertest: card string must be two characters")
    }

    rank, ok := rankMap[s[:1]]
    if !ok {
        panic("jokertest: rank not found")
    }

    suit, ok := suitMap[s[1:]]
    if !ok {
        panic("jokertest: suit not found")
    }

    for _, c := range hand.Cards() {
        if rank == c.Rank() && suit == c.Suit() {
            return c
        }
    }
    panic("jokertest: card not found")
}

var (
    rankMap = map[string] hand.Rank {
        "A": hand.Ace,
        "K": hand.King,
        "Q": hand.Queen,
        "J": hand.Jack,
        "T": hand.Ten,
        "9": hand.Nine,
        "8": hand.Eight,
        "7": hand.Seven,
        "6": hand.Six,
        "5": hand.Five,
        "4": hand.Four,
        "3": hand.Three,
        "2": hand.Two,
    }

    suitMap = map[string] hand.Suit {
        "s": hand.Spades,
        "h": hand.Hearts,
        "d": hand.Diamonds,
        "c": hand.Clubs,
    }
)
