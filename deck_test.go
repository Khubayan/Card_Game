package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	const expectedNumber int = 16
	d := newDeck()

	if len(d) != expectedNumber {
		t.Errorf("Expected deck length of %v, but got %v", expectedNumber, len(d))
	}
}