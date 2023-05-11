package internal

// Task represents a task table.
// ID is the unique identifier for the task.
// Title is the title of the task.
// Description provides additional details about the task.
// Status represents the current status of the task.
// CreatedAt stores the timestamp of when the task was created.
// UpdatedAt stores the timestamp of when the task was last updated.
type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string
	CreatedAt   int64
	UpdatedAt   int64
}
