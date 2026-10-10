package del

import (
	"context"
	"fmt"

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

func (r *Repository) DeleteTask(ctx context.Context, id int) error {
	tag, err := r.db.Exec(ctx,
		`DELETE FROM tasks WHERE id = $1`,
		id)
	if err != nil {
		return fmt.Errorf("Не удалось удалить задачу: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return apperrors.ErrTaskNotFound
	}

	return nil
}
