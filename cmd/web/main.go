package main

import (
	"log"
	"net/http"

	"github.com/ub2013210024/hello/handlers"
)

func main() {
	// Create a multiplexer
	mux := http.NewServeMux()
	// Create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))
	mux.HandleFunc("/greeting", handlers.Greeting)
	// ->home
	mux.HandleFunc("/", handlers.Home)
	// ->about
	mux.HandleFunc("/about", handlers.About)
	// Create our server
	mux.HandleFunc("/message/create", handlers.MessageCreate)
	// Create our server
	log.Print("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
