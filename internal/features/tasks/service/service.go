package tasks_service

import "github.com/shitaiv1ck/todolist/internal/core/domains"

type TasksService struct {
	rep TasksRepository
}

type TasksRepository interface {
	CreateTask(task *domains.Task) (*domains.Task, error)
}

func NewService(rep TasksRepository) *TasksService {
	return &TasksService{
		rep: rep,
	}
}

func (s *TasksService) CreateTask(task *domains.Task) (*domains.Task, error) {
	createdTask, err := s.rep.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
