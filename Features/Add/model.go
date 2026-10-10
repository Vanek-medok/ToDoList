package Add

import "time"

type TaskModel struct {
	ID           int
	Title        string
	Description  string
	Created_at   time.Time
	Completed    bool
	Completed_at *time.Time
}
