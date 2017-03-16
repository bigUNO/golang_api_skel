package main

import (
	"fmt"
)

var currentId int

var items Items
var neoitems neoItems

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
func createItem(name string) neoItem {
	newNeoItem := neoItem{}
	newNeoItem.Id = createXid()
	newNeoItem.Name = name
	neoitems = append(neoitems, newNeoItem)
	return newNeoItem
}

// findItem: lookup item by id (string)
func findItem(id string) neoItem {
	for _, nI := range neoitems {
	  found := getXidString(nI)
		if id == found {
			return nI
		}
	}
	// return empty item if not found
	return neoItem{}
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

func RepoDestroyItem(id int) error {
	for i, I := range items {
		if I.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find item with id of %d to delete", id)
}
