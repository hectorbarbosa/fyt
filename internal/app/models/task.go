package models

import "time"

type Task struct {
	Id          int64
	ProjectId   int64
	Title       string
	Description string
	DueDate     time.Time
	Doer        int32
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
