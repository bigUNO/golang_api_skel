package main

import (
	"testing"
	"unicode/utf8"
)

func TestCreateXidReturnsCorrectLength(t *testing.T) {
	expected := 20
	testID := createXid()

	actual := utf8.RuneCountInString(testID.String())

	if actual != expected {
		t.Errorf("got a sring of the wrong length: got %v want %v",
			actual, expected)
	}
}
