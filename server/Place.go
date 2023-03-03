package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PlaceFile struct {
	layers []struct {
		image   string
		objects []struct {
			x          int
			y          int
			width      int
			height     int
			properties []struct {
				name  string
				value string
			}
		}
	}
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

func LoadPlace(placename string) (*Place, error) {
	placepath := fmt.Sprintf("assets/%s.tmj", placename)
	r, err := os.Open(placepath)
	if err != nil {
		return nil, err
	}
	placefile := &PlaceFile{}
	place := &Place{}
	json.NewDecoder(r).Decode(placefile)
	for _, layer := range placefile.layers {
		if layer.image != "" {
			place.image = layer.image
		} else if layer.objects != nil {
			for _, object := range layer.objects {
				exit := Exit{
					rect: Rect{
						x:      object.x,
						y:      object.y,
						width:  object.width,
						height: object.height,
					},
				}
				for _, p := range object.properties {
					if p.name == "NextPlace" {
						exit.place = p.value
					}
				}
				place.exits = append(place.exits, exit)
			}
		}
	}
	return place, nil
}
