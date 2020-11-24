package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expectedDeckLength := 52
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "Two of Clubs"

	d := newDeck()
	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected %v, but got %v", expectedFirstCard, d[0])
	}

	if d[len(d) - 1] != expectedLastCard {
		t.Errorf("Expected %v, but got %v", expectedLastCard, d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktest"
	expectedDeckLength := 52

	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, len(deck))
	}

	os.Remove(fileName)

}

