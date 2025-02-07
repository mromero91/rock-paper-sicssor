package main

import (
	"log"
	"net/http"
	"rps/cmd/handlers"
)

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("cmd/static/"))

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Print("Server started on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
	http.ListenAndServe(port, nil)
}
