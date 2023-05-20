package taskv1

import "taskmanager/models"

// TaskRequestModel represents a Request Model for a Task.
type TaskRequestModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TaskResponseModel represents a Response Model for a Task.
type TaskResponseModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      models.TaskStatus
}
