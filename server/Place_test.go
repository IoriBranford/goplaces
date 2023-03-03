package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPlace(t *testing.T) {
	place, err := LoadPlace("assets/Train_Day.tmj")
	assert.NotNil(t, place)
	assert.Nil(t, err)
	assert.Contains(t, place.image, "Train_Day")
}
