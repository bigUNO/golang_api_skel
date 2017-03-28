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
	Route{"ItemIndex", "GET", "/v0/items", itemIndex},
	Route{"ItemCreate", "POST", "/v0/items", itemCreate},
	Route{"ItemShow", "GET", "/v0/items/{ItemId}", itemShow},
	Route{"ItemDelete", "DELETE", "/v0/items/{ItemId}", itemDelete},
}
