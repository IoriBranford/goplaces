package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PlaceFile struct {
	Layers []struct {
		Image   string `json:"image"`
		Objects []struct {
			X          int `json:"x"`
			Y          int `json:"y"`
			Width      int `json:"width"`
			Height     int `json:"height"`
			Properties []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"properties"`
		} `json:"objects"`
	} `json:"layers"`
}

type Place struct {
	image string
	exits []Exit
}

type Rect struct {
	x      int
	y      int
	width  int
	height int
}

type Exit struct {
	rect  Rect
	place string
}

func LoadPlace(placepath string) (*Place, error) {
	dir := filepath.Dir(placepath)
	r, err := os.Open(placepath)
	if err != nil {
		return nil, err
	}
	placefile := &PlaceFile{}
	place := &Place{}
	json.NewDecoder(r).Decode(placefile)
	for _, layer := range placefile.Layers {
		if layer.Image != "" {
			place.image = filepath.Join(dir, layer.Image)
		} else if layer.Objects != nil {
			for _, object := range layer.Objects {
				exit := Exit{
					rect: Rect{
						x:      object.X,
						y:      object.Y,
						width:  object.Width,
						height: object.Height,
					},
				}
				for _, p := range object.Properties {
					if p.Name == "NextPlace" {
						exit.place = p.Value
					}
				}
				place.exits = append(place.exits, exit)
			}
		}
	}
	return place, nil
}
