package main

import (
	"github.com/rs/xid"
)

// createXiD: Create new UUID
func createXid() xid.ID {
	guid := xid.New()
	return guid
}
