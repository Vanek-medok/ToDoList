package list

import (
	"context"
	"log/slog"
)

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
) ([]TaskDTO, error) {
	var tasksDTO []TaskDTO

	tasksModel, err := s.repository.TasksList(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Ошибка при получении списка задач", "error", err)
	}
	for _, task := range tasksModel {
		tasksDTO = append(tasksDTO, ToTaskDTO(&task))
	}
	return tasksDTO, nil
}
