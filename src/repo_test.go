package main

import (
  "testing"
  "fmt"
)

// Test creating an Item
func TestCreateItem(t *testing.T) {
	expected := "flowingrobes"
	testItem := createItem(expected)

	actual := getXidName(testItem)
	if actual != expected {
		t.Errorf("got a name mismatch: got %v wanted %v",
			actual, expected)
		fmt.Printf("%+v\n", testItem)
	}
}

// Test finding an Item
func TestfindItem(t *testing.T) {
	expected := "ghostfart"
  sniff := findItem(expected)

  actual := getXidName(sniff)
	if actual != expected {
		t.Errorf("cannot find item: got %v wanted %v",
			actual, expected)
	}
}
