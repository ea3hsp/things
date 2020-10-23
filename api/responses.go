package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ea3hsp/test/models"
)

// CreateThingsResponse encodes http create thing response
func CreateThingsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// payload
	payload := response.([]byte)
	res := map[string]interface{}{
		"response": string(payload),
	}
	// return
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(&res)
}

// ReadThingResponse encodes http read thing response
func ReadThingResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// payload
	thing := response.(*models.Thing)
	// return
	return json.NewEncoder(w).Encode(&thing)
}

// UpdateThingResponse encodes http update thing response
func UpdateThingResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// payload
	thing := response.(*models.Thing)
	// return
	return json.NewEncoder(w).Encode(&thing)
}
