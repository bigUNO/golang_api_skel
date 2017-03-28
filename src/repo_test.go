package main

import (
	"testing"
	"time"
	"fmt"

	"github.com/rs/xid"
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

// Test marking an item as deleted
func TestDeleteItem(t *testing.T) {
	id := xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}
	exampleTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testItem := Item{Id: id, Name: "beansprouts", Modified: exampleTime}
	expected := true

	testItem.SetDeleted()
	actual := getIsDeleted(testItem)
	if actual != true {
		t.Errorf("item has risen from the grave: got %v wanted %v",
		  actual, expected)
	}
}
