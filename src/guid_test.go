package main

import (
	"testing"
  "unicode/utf8"
)

func TestGetXidReturnsCorrectLength(t *testing.T) {
  expected := 20
  actual := utf8.RuneCountInString(getXid())

  if actual != expected {
    t.Errorf("got a sring of the wrong length: got %v want %v",
      actual, expected)
  }
}
