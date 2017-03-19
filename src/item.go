package main

import (
	"time"

	"github.com/rs/xid"
)

type neoItem struct {
	Id   xid.ID `json:"xid"`
	Name string `json:"name"`
}

type neoItems []neoItem

// SetDefaults: sets the UUID on creation
func (self *neoItem) SetDefaults() {
	self.Id = createXid()
}

// METHODS

// getXidString: return neoItems UUID string
func getXidString(nI neoItem) string {
	s := nI.Id
	return s.String()
}

// getXidName: return neoItems Name string
func getXidName(nI neoItem) string {
	return nI.Name
}

// getXidTimestamp: Returns timestamp of xid
func getXidTimestamp(nI neoItem) time.Time {
	t := nI.Id
	return t.Time()
}

// getXidCounter: Returns counter of xid
func getXidCounter(nI neoItem) int32 {
	c := nI.Id
	return c.Counter()
}
