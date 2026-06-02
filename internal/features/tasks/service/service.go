package tasks_service

import "github.com/shitaiv1ck/todolist/internal/core/domains"

type TasksService struct {
	rep TasksRepository
}

type TasksRepository interface {
	CreateTask(task *domains.Task) (*domains.Task, error)
	UpdateTask(task *domains.Task) (*domains.Task, error)
	FindByUserID(id int, userID int) (*domains.Task, error)
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

func (s *TasksService) UpdateTask(patch *domains.TaskPatch) (*domains.Task, error) {
	foundTask, err := s.rep.FindByUserID(patch.ID, patch.UserID)
	if err != nil {
		return nil, err
	}

	if err := foundTask.ApplyPatch(patch); err != nil {
		return nil, err
	}

	patchedTask, err := s.rep.UpdateTask(foundTask)
	if err != nil {
		return nil, err
	}

	return patchedTask, nil
}
