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
	Route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		HealthCheck,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
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
