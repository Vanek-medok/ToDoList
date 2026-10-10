package Add

import (
	"context"
	"errors"
	"fmt"
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
		return nil, errors.New("пустой заголовок")
	}

	taskModel := TaskModel{
		Title:       dto.Title,
		Description: dto.Description,
	}

	task, err := s.repository.CreateTask(ctx, &taskModel)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании задачи:%w", err)
	}

	return task, nil
}
