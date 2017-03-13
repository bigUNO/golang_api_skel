package main

import (
	"fmt"
)

var currentId int

var items Items
var testItems neoItems

func init() {
	// creating test items
	RepoCreateItem(Item{Name: "candlestick"})
	RepoCreateItem(Item{Name: "rope"})
	RepoCreateItem(Item{Name: "lead pipe"})
	RepoCreateItem(Item{Name: "revolver"})
	RepoCreateItem(Item{Name: "wrench"})

}

func RepoFindItem(id int) Item {
	for _, I := range items {
		if I.Id == id {
			return I
		}
	}
	// return empty item if not found
	return Item{}
}

func RepoCreateItem(I Item) Item {
	currentId += 1
	I.Id = currentId
	items = append(items, I)
	return I
}

// Replacement item creater
func CreateTestItem(nI neoItem) neoItem {
	testItems = append(testItems, nI)
	return nI
}

func RepoDestroyItem(id int) error {
	for i, I := range items {
		if I.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find item with id of %d to delete", id)
}
