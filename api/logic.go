package api

import (
	"context"

	log "github.com/go-kit/kit/log"

	"github.com/ea3hsp/test/models"
)

type service struct {
	ctx    context.Context
	repo   Repository
	logger log.Logger
}

// NewService creates new service
func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateThing(ctx context.Context, thing ...models.Thing) (interface{}, error) {
	// create things
	err := s.repo.CreateThing(ctx, thing...)
	if err != nil {
		return nil, err
	}
	return []byte("things created successfully"), nil
}

func (s *service) ReadThing(ctx context.Context, key string) (interface{}, error) {
	thing, err := s.repo.ReadThing(ctx, key)
	if err != nil {
		return nil, err
	}
	return thing, nil
}

func (s *service) DeleteThing(ctx context.Context, key string) (interface{}, error) {
	s.repo.DeleteThing(ctx, key)
	return nil, nil
}

func (s *service) UpdateThing(ctx context.Context, thing models.Thing) (interface{}, error) {
	s.repo.UpdateThing(ctx, thing)
	return &thing, nil
}
