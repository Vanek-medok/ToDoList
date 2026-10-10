package list

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

func (r *Repository) TasksList(ctx context.Context) ([]TaskModel, error) {
	sqlQuery :=
		`
        SELECT
            id,
            title,
            description,
            created_at,
            completed,
            completed_at
        FROM tasks
        ORDER BY id;
    `

	rows, err := r.db.Query(ctx, sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения задач: %w", err)
	}

	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Created_at,
			&task.Completed,
			&task.Completed_at,
		)
		if err != nil {
			return nil, fmt.Errorf("Ошибка чтения задачи: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка обхода результатов: %w", err)
	}

	return tasks, nil
}
