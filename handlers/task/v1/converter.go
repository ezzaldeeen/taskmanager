package taskv1

import (
	"github.com/google/uuid"
	"taskmanager/models"
	"time"
)

// fromIDOtoDTO converts an IDO (Intermediate Data Object) to
// a DTO (Data Transfer Object) for a rest.
// It creates a new DTO instance and initializes its fields
// with the corresponding values from the IDO.
func fromIDOtoDTO(ido models.TaskIDO) TaskDTO {
	return TaskDTO{
		Title:       ido.Title,
		Description: ido.Description,
	}
}

// fromIDOtoDTO converts a DTO (Data Transfer Object) to
// an IDO (Intermediate Data Object) for a rest.
// It creates a new D instance and initializes its fields
// with the corresponding values from the DTO.
func fromDTOtoIDO(dto TaskDTO) models.TaskIDO {
	return models.TaskIDO{
		Id:          uuid.NewString(),
		Title:       dto.Title,
		Description: dto.Description,
		Status:      models.TaskActiveStatus,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
}

// fromIDOsToDTOs converts DTOs (Data Transfer Object) to
// IDOs (Intermediate Data Object) for a rest.
// It creates a new DTO instances and initializes their fields
// with the corresponding values from the IDOs.
func fromIDOsToDTOs(idos []models.TaskIDO) []TaskDTO {
	var dtos []TaskDTO
	for _, ido := range idos {
		dto := TaskDTO{
			Title:       ido.Title,
			Description: ido.Description,
		}
		dtos = append(dtos, dto)
	}
	return dtos
}
