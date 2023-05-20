package tasksvc

// TaskRepo wrapper of the DB accessing logic
type TaskRepo interface {
	TaskCreatorRepo
	TasksGetterRepo
}

// Service represents a service responsible
// for handling business logic related to tasks.
type Service struct {
	repo TaskRepo
}

// NewService creates a new instance of the Service.
func NewService(repo TaskRepo) Service {
	return Service{
		repo: repo,
	}
}
