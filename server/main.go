package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	world := NewWorld()
	world.LoadPlaces("assets/*.tmj")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	})
	r.HandleFunc("/goto", func(w http.ResponseWriter, r *http.Request) {
		world.handleGoTo(w, r)
	})

	fileServer := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	log.Fatal(http.ListenAndServe(":8080", r))
}
