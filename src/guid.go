package main

import (
	"github.com/rs/xid"
)

func getXid() (string) {
  guid := xid.New()
  return guid.String()
}

func getXidCounter(x string) int32 {
	guid := xid.New()
  counter := guid.Counter(x)
	return counter
}
