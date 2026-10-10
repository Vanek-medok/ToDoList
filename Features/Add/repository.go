package Add

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateTask(ctx context.Context, task *TaskModel) (*TaskModel, error) {
	sqlQuery := `
	INSERT INTO tasks(title, description, completed, completed_at)
	VALUES($1, $2, $3, $4)
	RETURNING id, title, description, created_at, completed, completed_at;
	`
	var NewTask TaskModel

	err := r.db.QueryRow(ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.Completed_at,
	).Scan(
		&NewTask.ID,
		&NewTask.Title,
		&NewTask.Description,
		&NewTask.Created_at,
		&NewTask.Completed,
		&NewTask.Completed_at,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при записи или сканировании задачи: %w", err)
	}

	return &NewTask, nil
}
