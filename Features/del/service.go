package del

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

func (s *Service) DeleteTask(ctx context.Context, id int) error {

	if id <= 0 {
		return apperrors.ErrInvalidTaskID
	}

	return s.repository.DeleteTask(ctx, id)
}
