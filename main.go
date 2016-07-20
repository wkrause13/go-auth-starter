package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"go-auth-starter/middleware"
)

func DemoIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	fmt.Fprintf(w, "Hi there, I love %s!", "Demos")
	return
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "UPDATE", "OPTIONS"},
	})

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middleware.AddContentTypeHeader))
	n.Use(c)
	router := NewRouter()

	router.Handle("GET", "/", DemoIndexHandler)

	n.UseHandler(router)
	n.Run("0.0.0.0:8080")

}
