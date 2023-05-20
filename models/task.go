package models

type Status int

const (
	TaskActiveStatus    Status = 1
	TaskCompletedStatus Status = 0
	TaskDeletedStatus   Status = -1
)

// TaskIDO represents an intermediate Data Object for a rest.
type TaskIDO struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	// todo: add reference for the author
}
