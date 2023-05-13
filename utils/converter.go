package utils

import "taskmanager/models/task"

func ConvertIDOStoDTOS(idos []task.IDO) []task.DTO {
	dtos := make([]task.DTO, len(idos))
	for i, ido := range idos {
		dtos[i] = *ido.ToDTO()
	}
	return dtos
}
