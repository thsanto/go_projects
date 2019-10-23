package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Errou... deveria ser %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Devia ser Ace of Spades, mas %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Devia ser Four of Clubs, but... %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16, but... %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
