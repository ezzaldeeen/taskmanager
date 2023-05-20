package tasksvc

type TaskRepo interface {
	TaskGetterRepo
	TaskCreatorRepo
}

type Service struct {
	repo TaskRepo
}

func NewService(repo TaskRepo) Service {
	return Service{
		repo: repo,
	}
}
