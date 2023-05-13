package tasksvc

import (
	"taskmanager/models/task"
)

type Repo interface {
	Create(ido *task.IDO) error
	GetAll() ([]task.IDO, error)
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) CreateTask(ido *task.IDO) error {
	err := s.repo.Create(ido)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetAllTask() ([]task.IDO, error) {
	idos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return idos, nil
}
