package main

import (
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// placeholder for incoming items
type tempItem struct {
	Name string `json:"name"`
}

// Index: delete this later?
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

// healthCheck: basic health-check
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

// itemIndex: return all items in index
func itemIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

// itemCreate: add a new item (via createItem) to the index
func itemCreate(w http.ResponseWriter, res *http.Request) {
	var ItemName string

	body, err := ioutil.ReadAll(io.LimitReader(res.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := res.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &ItemName); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
  // unwrap JSON key:value
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	var tempItemName = result["name"].(string)
	I := createItem(tempItemName)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(I); err != nil {
		panic(err)
	}
}

// itemShow: Return all items in index
func itemShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemXid string
	var err bool
	if itemXid, err = vars["ItemId"]; err != true {
		panic(err)
	}

	item := findItem(itemXid)
	c := item.Id
	if c.Counter() > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(item); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
