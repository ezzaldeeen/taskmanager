package task

type Status int

const (
	Active    Status = 1
	Completed Status = 0
	Deleted   Status = -1
)

// IDO represents an intermediate Data Object for a task.
type IDO struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ToDTO converts an IDO (Intermediate Data Object) to
// a DTO (Data Transfer Object). It creates and returns a new instance of DTO
// with the Title and Description fields
// The remaining fields in the DTO will be set to their default values.
func (t *IDO) ToDTO() *DTO {
	return &DTO{
		Title:       t.Title,
		Description: t.Description,
	}
}
