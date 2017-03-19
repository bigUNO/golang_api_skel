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
	Route{"ItemIndex", "GET", "/items", itemIndex},
	Route{"ItemCreate", "POST", "/items", itemCreate},
	Route{"ItemShow", "GET", "/items/{ItemId}", itemShow},
	Route{"ItemDelete", "DELETE", "/items/{ItemId}", itemDelete},
}
