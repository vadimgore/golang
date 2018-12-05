package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Invalid number of cards (%v) in a new deck", len(d))
	}
}

func TestFileSave(t *testing.T) {
	d1 := newDeck()
	d1.shuffle()
	s1 := d1.toString()
	d1.saveToFile("MyCards")

	d2 := newDeckFromFile("MyCards")
	s2 := d2.toString()

	os.Remove("MyCards")

	if s1 != s2 {
		t.Errorf("Saving and restoring deck has failed")
	}
}
