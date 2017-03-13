package main

import (
	"github.com/rs/xid"
)

func createXid() xid.ID {
	guid := xid.New()
	return guid
}
