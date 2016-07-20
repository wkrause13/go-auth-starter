package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		pattern := "/" + route.Pattern
		router.Handle(route.Method, pattern, route.Handle)
	}
	return router
}
