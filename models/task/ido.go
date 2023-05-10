package task

import (
	"github.com/google/uuid"
	"time"
)

type DTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ToIDO converts a DTO (Data Transfer Object) to
// an IDO (Intermediate Data Object) for a task.
// It creates a new IDO instance and initializes its fields
// with the corresponding values from the DTO.
func (t *DTO) ToIDO() *IDO {
	return &IDO{
		Id:          uuid.NewString(),
		Title:       t.Title,
		Description: t.Description,
		Status:      Active,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}
