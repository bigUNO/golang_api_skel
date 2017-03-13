package main

import (
	"testing"
	"unicode/utf8"
)


func TestGetId(t *testing.T) {
	expected := 20
	testID := createXid()

	actual := utf8.RuneCountInString(testID.String())
	if actual != expected {
		t.Errorf("got a sring of the wrong length: got %v wanted %v",
			actual, expected)
	}
}

/*
func TestIDString(t *testing.T) {

	id := ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}
  expected := "9m4e2mr0ui3e8a215n4g"
	actual := id.String()

	//assert.Equal(t, "9m4e2mr0ui3e8a215n4g", id.String())
}*/
