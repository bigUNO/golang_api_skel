package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/xid"
)

func setup() {
	id := xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}
	exampleTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	waldo := Item{Id: id, Name: "ghostfart", Modified: exampleTime}
	fmt.Printf("%v", waldo)
}

func TestMain(m *testing.M) {
	setup()
}
