package api

import (
	"context"
	"errors"

	"github.com/ea3hsp/test/models"
)

var (
	// ErrorRepoConnect ...
	ErrorRepoConnect = errors.New("database connection error")
	// ErrorRepoCreate ...
	ErrorRepoCreate = errors.New("thing create error")
	// ErrorAffectedRowsExpected ...
	ErrorAffectedRowsExpected = errors.New("affected rows expected")
	// ErrorThingInsert ...
	ErrorThingInsert = errors.New("insert thing error")
	// ErrorThingRead ...
	ErrorThingRead = errors.New("read thing error")
)

// Repository represents repository
type Repository interface {
	CreateThing(ctx context.Context, things ...models.Thing) error
	ReadThing(ctx context.Context, key string) (*models.Thing, error)
	UpdateThing(ctx context.Context, thing models.Thing) error
	DeleteThing(ctx context.Context, key string) error
	Close(ctx context.Context) error
}
