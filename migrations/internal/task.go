package internal

import "taskmanager/models"

// Task represents a rest table.
// ID is the unique identifier for the rest.
// Title is the title of the rest.
// Description provides additional details about the rest.
// Status represents the current status of the rest.
// CreatedAt stores the timestamp of when the rest was created.
// UpdatedAt stores the timestamp of when the rest was last updated.
type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      models.TaskStatus
	CreatedAt   int64
	UpdatedAt   int64
	// todo: add reference for the author
}
