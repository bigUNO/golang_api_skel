package main

import (
	"reflect"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/rs/xid"
)

// Test mocking up xid from byte code
func TestXid(t *testing.T) {
	id := xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}
	expected := "9m4e2mr0ui3e8a215n4g"
	actual := id.String()
	if actual != expected {
		t.Errorf("got %v wanted %v",
			actual, expected)
	}
}

// Test creating a new UUID
func TestCreateXid(t *testing.T) {
	expected := 20
	testID := createXid()

	actual := utf8.RuneCountInString(testID.String())
	if actual != expected {
		t.Errorf("got a string of the wrong length: got %v wanted %v",
			actual, expected)
	}
}

// Test ability to get string version of UUID
func TestGetXidString(t *testing.T) {
	id := xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}
	testItem := neoItem{Id: id, Name: "beansprouts"}

	expected := "string"
	s := getXidString(testItem)
	actual := reflect.TypeOf(s).Kind()

	if actual != reflect.String {
		t.Errorf("got the wrong type: got %v wanted %v",
			actual, expected)
	}
}

func TestGetXidTimestamp(t *testing.T) {
	id := xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}
	testItem := neoItem{Id: id, Name: "beansprouts"}
	currentTime := time.Now()

	//expected := 30
	actual := getXidTimestamp(testItem)
	expected := actual.Sub(currentTime)
	if expected > 1 {
		t.Errorf("fucked up: got %v wanted %v",
			actual, expected)
	}
}
