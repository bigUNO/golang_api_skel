package main

import (
	"testing"
  "unicode/utf8"
)

func TestGetId(t *testing.T) {
  expected := 20

  actual := utf8.RuneCountInString(getXid())
  if actual != expected {
    t.Errorf("got a sring of the wrong length: got %v wanted %v",
      actual, expected)
  }
}

func TestItemCounter(t *testing.T) {
  //testXid := CreateTestItem(neoItem{Name: "rope"})
  testXid := "9m4e2mr0ui3e8a215n4g"
  expected := 2147483647

  actual := getXidCounter(testXid)
  if actual != expected {
    t.Errorf("got a different counter: got %v wanted %v", actual, expected)
  }
}
