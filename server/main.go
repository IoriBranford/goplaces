package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Place struct {
	Id    int
	Image string
	Exits []Exit
}

type Exit struct {
	ClickRect struct {
		x      int
		y      int
		width  int
		height int
	}
	NextPlace Place
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
