package taskv1

// TaskDTO represents a Data Transfer Object for a rest.
type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
