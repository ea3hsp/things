package api

import (
	"context"

	"github.com/ea3hsp/test/models"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints definitions
type Endpoints struct {
	CreateThingEP endpoint.Endpoint
	ReadThingEP   endpoint.Endpoint
	UpdateThingEP endpoint.Endpoint
	DeleteThingEP endpoint.Endpoint
}

// MakeEndpoints creates end points defined
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateThingEP: func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.([]models.Thing)
			return s.CreateThing(ctx, req...)
		},
		ReadThingEP: func(ctx context.Context, request interface{}) (interface{}, error) {
			id := request.(string)
			return s.ReadThing(ctx, id)
		},
		UpdateThingEP: func(ctx context.Context, request interface{}) (interface{}, error) {
			thing := request.(models.Thing)
			return s.UpdateThing(ctx, thing)
		},
		DeleteThingEP: func(ctx context.Context, request interface{}) (interface{}, error) {
			id := request.(string)
			return s.DeleteThing(ctx, id)
		},
	}
}
