package taskv1

import (
	"github.com/google/uuid"
	"taskmanager/models"
	"time"
)

// convertToResponseModel converts an IDO (Intermediate Data Object) to
// a DTO (Data Transfer Object) for a rest.
// It creates a new DTO instance and initializes its fields
// with the corresponding values from the IDO.
func convertToResponseModel(ido models.TaskIDO) TaskResponseModel {
	return TaskResponseModel{
		Title:       ido.Title,
		Description: ido.Description,
		Status:      ido.Status,
	}
}

// convertToTaskIDO converts a DTO (Data Transfer Object) to
// an IDO (Intermediate Data Object) for a rest.
// It creates a new D instance and initializes its fields
// with the corresponding values from the DTO.
func convertToTaskIDO(dto TaskRequestModel) models.TaskIDO {
	return models.TaskIDO{
		Id:          uuid.NewString(),
		Title:       dto.Title,
		Description: dto.Description,
		Status:      models.TaskActiveStatus,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}

// convertTasksToResponseModels converts DTOs (Data Transfer Object) to
// IDOs (Intermediate Data Object) for a rest.
// It creates a new DTO instances and initializes their fields
// with the corresponding values from the IDOs.
func convertTasksToResponseModels(idos []models.TaskIDO) []TaskResponseModel {
	var dtos []TaskResponseModel
	for _, ido := range idos {
		dto := TaskResponseModel{
			Title:       ido.Title,
			Description: ido.Description,
			Status:      ido.Status,
		}
		dtos = append(dtos, dto)
	}
	return dtos
}
