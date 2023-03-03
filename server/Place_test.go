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
	assert.Equal(t, place.image, "image.png")
	assert.Greater(t, len(place.exits), 0)
	assert.Equal(t, place.exits[0].rect.x, 10)
	assert.Equal(t, place.exits[0].rect.y, 20)
	assert.Equal(t, place.exits[0].rect.width, 30)
	assert.Equal(t, place.exits[0].rect.height, 40)
	assert.Equal(t, place.exits[0].place, "nextplace")
}
