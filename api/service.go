package api

import (
	"context"

	"github.com/ea3hsp/test/models"
)

// Service represents service
type Service interface {
	CreateThing(ctx context.Context, things ...models.Thing) (interface{}, error)
	ReadThing(ctx context.Context, id string) (interface{}, error)
	UpdateThing(ctx context.Context, thing models.Thing) (interface{}, error)
	DeleteThing(ctx context.Context, id string) (interface{}, error)
}
