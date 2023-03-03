package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPlace(t *testing.T) {
	place, err := LoadPlace("Train_Day")
	assert.NotNil(t, place)
	assert.Nil(t, err)
	assert.Contains(t, place.image, "Train_Day")
}
