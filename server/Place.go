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
	Name  string `json:"name"`
	Image string `json:"image"`
	Exits []Exit `json:"exits"`
}

type Exit struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Place  string `json:"place"`
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
	place.Name = strings.TrimSuffix(filepath.Base(placepath), filepath.Ext(placepath))
	place.Image = filepath.Join(filepath.Dir(placepath), place.Image)
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
			place.Image = layer.Image
		} else if layer.Objects != nil {
			for _, object := range layer.Objects {
				exit := Exit{
					X:      object.X,
					Y:      object.Y,
					Width:  object.Width,
					Height: object.Height,
				}
				for _, p := range object.Properties {
					if p.Name == "NextPlace" {
						exit.Place = p.Value
					}
				}
				place.Exits = append(place.Exits, exit)
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
		places[place.Name] = place
	}
	return places
}
