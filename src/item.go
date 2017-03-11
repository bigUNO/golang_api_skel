package main

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type neoItem struct {
	Id   string    `json:"id"`
	Name string `json:"name"`
}

type Items []Item
type neoItems []neoItem

func (self *neoItem) SetDefaults() {
   self.Id = getXid()
}

// NewSomething create new instance of Something
func NewNeoItem(name string) neoItem {
   newNeoItem := neoItem{}
   newNeoItem.Id = getXid()
   return newNeoItem
}

// Name receives a copy of Foo since it doesn't need to modify it.
func (nI neoItem) getId() string {
    return nI.Id
}
/*
func (xid string) getItemCounter(xid string) int32 {
	return guid.Counter(xid)
}
*/
