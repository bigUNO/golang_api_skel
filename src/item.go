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

// NewNeoItem: NewSomething create new instance of Something
func NewNeoItem(name string) neoItem {
	newNeoItem := neoItem{}
	newNeoItem.Id = createXid()
	return newNeoItem
}

// getXidString: return neoItems UUID string
func getXidString(nI neoItem) string {
	s := nI.Id
	return s.String()
}

// getXidTimestamp: Returns timestamp of xid
func getXidTimestamp(nI neoItem) time.Time {
	t := nI.Id
	return t.Time()
}
