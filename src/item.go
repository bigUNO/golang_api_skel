package main

import (
	"time"

	"github.com/rs/xid"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type neoItem struct {
	Id   xid.ID `json:"xid"`
	Name string `json:"name"`
}

type Items []Item
type neoItems []neoItem

func (self *neoItem) SetDefaults() {
	self.Id = createXid()
}

// NewSomething create new instance of Something
func NewNeoItem(name string) neoItem {
	newNeoItem := neoItem{}
	newNeoItem.Id = createXid()
	return newNeoItem
}

// Name receives a copy of Foo since it doesn't need to modify it.
func (nI neoItem) getXid() xid.ID {
	return nI.Id
}

func (nI neoItem) getXidString() string {
	s := nI.Id
	return s.String()
}

func (nI neoItem) getXidTimestamp() time.Time {
	t := nI.Id
	return t.Time()
}
