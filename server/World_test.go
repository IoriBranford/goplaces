package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorldGoTo(t *testing.T) {
	world := NewWorld()
	world.places["Place1"] = &Place{
		Name: "Place1",
	}
	req := httptest.NewRequest(http.MethodGet, "/goto", strings.NewReader(`{"place":"Place1"}`))
	w := httptest.NewRecorder()
	world.handleGoTo(w, req)
	res := w.Result()
	defer res.Body.Close()
	newplace := &Place{}
	err := json.NewDecoder(res.Body).Decode(newplace)
	assert.Nil(t, err)
	assert.Equal(t, "Place1", newplace.Name)
}
