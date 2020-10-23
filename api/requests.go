package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ea3hsp/test/models"
	"github.com/gorilla/mux"
)

// CreateThingsRequest decodes create things request
func CreateThingsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	things := []models.Thing{}
	// things body request
	err := json.NewDecoder(r.Body).Decode(&things)
	if err != nil {
		return nil, err
	}
	return things, nil
}

// ReadThingRequest decodes read thing request
func ReadThingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return vars["key"], nil
}

// UpdateThingRequest ...
func UpdateThingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	thing := models.Thing{}
	// things body request
	err := json.NewDecoder(r.Body).Decode(&thing)
	if err != nil {
		return nil, err
	}
	return thing, nil
}
