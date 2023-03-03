package main

import (
	"encoding/json"
	"net/http"
)

type World struct {
	places map[string]*Place
}

func NewWorld() *World {
	return &World{
		places: make(map[string]*Place),
	}
}

func (w *World) LoadPlaces(pattern string) {
	w.places = LoadPlaces("assets/*.tmj")
}

func (w *World) handleGoTo(writer http.ResponseWriter, reader *http.Request) {
	goTo := &struct {
		Place string `json:"place"`
	}{}
	json.NewDecoder(reader.Body).Decode(goTo)

	place := w.places[goTo.Place]
	if place == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(writer).Encode(place)
}
