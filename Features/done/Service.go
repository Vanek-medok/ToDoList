package done

import (
	"context"

	"main.go/core/apperrors"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) DoneTask(ctx context.Context, id int) (*TaskModel, error) {
	if id <= 0 {
		return nil, apperrors.ErrInvalidTaskID
	}

	return s.repository.DoneTask(ctx, id)
}
