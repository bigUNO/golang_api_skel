package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", index},
	Route{"HealthCheck", "GET", "/healthcheck", healthCheck},
	Route{"NeoItemIndex", "GET", "/neoitems", itemIndex},
	Route{"NeoItemCreate", "POST", "/neoitems", itemCreate},
	Route{"NeoItemShow", "GET", "/neoitems/{NeoItemId}", itemShow},
	Route{
		"ItemIndex",
		"GET",
		"/items",
		ItemIndex,
	},
	Route{
		"ItemCreate",
		"POST",
		"/items",
		ItemCreate,
	},
	Route{
		"ItemShow",
		"GET",
		"/items/{itemId}",
		ItemShow,
	},
}
