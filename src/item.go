package main

import (
	"time"

	"github.com/rs/xid"
)

type Item struct {
	Id        xid.ID    `json:"xid"`
	Name      string    `json:"name"`
	Modified  time.Time `json:"modified"`
	IsDeleted bool      `json:"isdeleted"`
}

type Items []Item

// SetModified: Sets last modified date
func (self *Item) SetModified() {
	self.Modified = time.Now()
	return
}

// SetDeleted: Marks the record as deleted
func (self *Item) SetDeleted() {
	self.IsDeleted = true
	return
}

// External METHODS

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

// getItemModified: Returns modified time
func getItemModified(I Item) time.Time {
	return I.Modified
}

// getIsDeleted: Returns is item is deleted or not
func getIsDeleted(I Item) bool {
	return I.IsDeleted
}
