package main

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Items []Item
