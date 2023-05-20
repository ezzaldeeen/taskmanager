package models

type TaskStatus string

const (
	TaskActiveStatus    TaskStatus = "active"
	TaskCompletedStatus TaskStatus = "completed"
	TaskDeletedStatus   TaskStatus = "deleted"
)

// TaskIDO represents an intermediate Data Object for a rest.
type TaskIDO struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   int64      `json:"updated_at"`
	// todo: add reference for the author
}
