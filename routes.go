package main

import (
	"go-auth-starter/handler"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"hello",
		handler.HelloWorld,
	},
}
