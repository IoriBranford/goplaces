package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	name  string
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
	r, err := os.Open(placepath)
	if err != nil {
		return nil, err
	}
	place, err := ReadPlace(r)
	if err != nil {
		return nil, err
	}
	place.name = strings.TrimSuffix(filepath.Base(placepath), filepath.Ext(placepath))
	place.image = filepath.Join(filepath.Dir(placepath), place.image)
	return place, nil
}

func ReadPlace(r io.Reader) (*Place, error) {
	placefile := &PlaceFile{}
	err := json.NewDecoder(r).Decode(placefile)
	if err != nil {
		return nil, err
	}

	place := &Place{}
	for _, layer := range placefile.Layers {
		if layer.Image != "" {
			place.image = layer.Image
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

func LoadPlaces(pattern string) map[string]*Place {
	places := map[string]*Place{}
	placefiles, _ := filepath.Glob(pattern)
	for _, placefile := range placefiles {
		place, err := LoadPlace(placefile)
		if err != nil {
			fmt.Println(err)
			continue
		}
		places[place.name] = place
	}
	return places
}
