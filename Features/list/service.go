package list

import "context"

type Service struct {
	repository *Repository
}

func NewService(repsitory *Repository) *Service {
	return &Service{
		repository: repsitory,
	}
}

func (s *Service) TasksList(
	ctx context.Context,
) ([]TaskModel, error) {
	return s.repository.TasksList(ctx)
}
