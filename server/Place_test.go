package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPlace(t *testing.T) {
	placejson := `{
		"layers": [
			{
				"image": "image.png"
			},
			{
				"objects": [
					{
						"x": 10,
						"y": 20,
						"width": 30,
						"height": 40,
						"properties": [
							{
								"name": "NextPlace",
								"value": "nextplace"
							}
						]
					}
				]
			}
		]
	}`
	place, err := ReadPlace(strings.NewReader(placejson))
	assert.NotNil(t, place)
	assert.Nil(t, err)
	assert.Equal(t, place.Image, "image.png")
	assert.Greater(t, len(place.Exits), 0)
	assert.Equal(t, place.Exits[0].X, 10)
	assert.Equal(t, place.Exits[0].Y, 20)
	assert.Equal(t, place.Exits[0].Width, 30)
	assert.Equal(t, place.Exits[0].Height, 40)
	assert.Equal(t, place.Exits[0].Place, "nextplace")
}
