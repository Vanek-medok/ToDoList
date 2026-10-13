package done

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/core/apperrors"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) DoneTask(ctx context.Context, id int) (*TaskModel, error) {
	sqlQuery := `
    UPDATE tasks
    SET completed = TRUE,
        completed_at = NOW()
    WHERE id = $1
      AND completed = FALSE
    RETURNING id, title, description, created_at, completed, completed_at;
	`
	var task TaskModel

	err := r.db.QueryRow(ctx, sqlQuery, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Created_at,
		&task.Completed,
		&task.Completed_at,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apperrors.ErrTaskNotFound
	}

	if err != nil {
		slog.ErrorContext(ctx, "не удалось выполнить задачу:", "error", err)
		return nil, fmt.Errorf("Не удалось выполнить задачу: %w", err)
	}

	return &task, nil
}
