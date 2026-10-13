package Add

import (
	"context"
	"errors"
	"log/slog"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) CreateTaskModel(ctx context.Context, dto *TaskDTO) (*TaskModel, error) {

	if dto.Title == "" {
		slog.ErrorContext(ctx, "Пустой заголовок", "error", errors.New("Пустой заголовок"))
		return nil, errors.New("пустой заголовок")
	}

	taskModel := TaskModel{
		Title:       dto.Title,
		Description: dto.Description,
	}

	task, err := s.repository.CreateTask(ctx, &taskModel)
	if err != nil {
		slog.ErrorContext(ctx, "Не удалось создать задачу:", "error", err)
		return nil, err
	}

	return task, nil
}
