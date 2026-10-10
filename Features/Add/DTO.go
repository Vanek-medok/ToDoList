package Add

import "time"

type TaskDTO struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Created_at   time.Time  `json:"created_at"`
	Completed    bool       `json:"completed"`
	Completed_at *time.Time `json:"completed_at"`
}

func ToTaskDTO(t TaskModel) TaskDTO {
	return TaskDTO{
		ID:           t.ID,
		Title:        t.Title,
		Description:  t.Description,
		Created_at:   t.Created_at,
		Completed:    t.Completed,
		Completed_at: t.Completed_at,
	}
}
