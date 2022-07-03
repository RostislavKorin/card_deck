package main

import (
	"testing"
	"os"
	"io/ioutil"
	"reflect"
)

func TestNewDeckSize(t *testing.T) {
	d := newDeck()
	deckLen := len(d)

	if deckLen != 16 {
		t.Errorf("Expected deck length of 16, but got %v", deckLen)
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected that first card is Spades of Ace, but got %v", d[0])
	}

	if d[len(d) - 1] != "Clubs of Four" {
		t.Errorf("Expected that last card is Clubs of Four, but got %v", d[len(d) - 1])
	}

	d.shuffle()

	if !reflect.DeepEqual(d, newDeck()) {
		t.Errorf("Original deck is equal to the shuffled deck")
	}
}


func TestDeckFile(t *testing.T) {
	filename := "tes_my_cards"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	_, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Errorf("Can't read the file")
	} 
	
	if !reflect.DeepEqual(newDeckFromFile(filename), d){
		t.Errorf("Deck from file is not eql to the given deck")
	}

	os.Remove(filename)
}

