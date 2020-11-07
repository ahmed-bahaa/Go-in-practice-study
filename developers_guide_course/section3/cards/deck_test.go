package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck length of 40, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck to start with Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected deck to end with Ten of Clubs, but got %v", d[len(d)-1])
	}
}

func TestWriteToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testfile")

	d := newDeck()
	d.writeToFile("_testfile")

	loadedDeck := newDeckFromFile("_testfile")
	if len(loadedDeck) != 40 {
		t.Errorf("Expected deck with length 40, but got %v", len(loadedDeck))
	}

	os.Remove("_testfile")
}
