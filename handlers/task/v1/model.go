package taskv1

import (
	"github.com/google/uuid"
	"taskmanager/models"
	"time"
)

// Task represents a Data Transfer Object for a rest.
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ToIDO converts a DTO (Data Transfer Object) to
// an IDO (Intermediate Data Object) for a rest.
// It creates a new IDO instance and initializes its fields
// with the corresponding values from the DTO.
func (t *Task) ToIDO() *models.IDO {
	return &models.IDO{
		Id:          uuid.NewString(),
		Title:       t.Title,
		Description: t.Description,
		Status:      models.Active,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}

// ToDTO converts an IDO (Intermediate Data Object) to
// todo: should be moved to the handler
// a DTO (Data Transfer Object). It creates and returns a new instance of DTO
// with the Title and Description fields
// The remaining fields in the DTO will be set to their default values.
func convertFromIDOtoTask(ido *models.IDO) Task {
	return Task{
		Title:       ido.Title,
		Description: ido.Description,
	}
}
