package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	//assert an element in a slice
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	deckName := "_decktesting"

	os.Remove(deckName)

	d := newDeck()
	d.saveToFile(deckName)

	loadedDeck, _ := readFile(deckName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	os.Remove(deckName)
}
