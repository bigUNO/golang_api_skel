package main

import (
	"time"

	"github.com/rs/xid"
)

type Item struct {
	Id   xid.ID `json:"xid"`
	Name string `json:"name"`
}

type Items []Item

// SetDefaults: sets the UUID on creation
func (self *Item) SetDefaults() {
	self.Id = createXid()
}

// METHODS

// getXidString: return Items UUID string
func getXidString(I Item) string {
	s := I.Id
	return s.String()
}

// getXidName: return Items Name string
func getXidName(I Item) string {
	return I.Name
}

// getXidTimestamp: Returns timestamp of xid
func getXidTimestamp(I Item) time.Time {
	t := I.Id
	return t.Time()
}

// getXidCounter: Returns counter of xid
func getXidCounter(I Item) int32 {
	c := I.Id
	return c.Counter()
}
