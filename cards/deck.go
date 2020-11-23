package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// create & return new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

// create a 'deal' function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// create a 'print()' method
// with a 'receiver' of type deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) shuffle() {

}