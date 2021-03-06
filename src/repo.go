package main

import (
	"fmt"
)

var items Items

func init() {
	// creating test items
	createItem("candlestick")
	createItem("rope")
	createItem("lead pipe")
	createItem("revolver")
	createItem("wrench")
	createItem("horseshoe")
}

// createItem: NewSomething create new instance of item
func createItem(name string) Item {
	newItem := Item{}
	newItem.Id = createXid()
	newItem.Name = name
	items = append(items, newItem)
	return newItem
}

// findItem: lookup item by id (string)
func findItem(id string) Item {
	for _, I := range items {
		found := getXidString(I)
		if id == found {
			return I
		}
	}
	// return empty item if not found
	return Item{}
}

// deleteItem: mark item as deleted
func deleteItem(id string) error {
	deathrow := findItem(id)
	for i, _ := range items {
		found := getXidString(deathrow)
		if id == found {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find item with id of %v to delete", id)
}
