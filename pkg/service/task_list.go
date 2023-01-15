package service

import (
	rsapi "github.com/wanna-beat-by-bit/rest-golang"
	"github.com/wanna-beat-by-bit/rest-golang/pkg/repository"
)

type TaskListService struct {
	repo repository.TaskList
}

func NewTaskListService(repo repository.TaskList) *TaskListService {
	return &TaskListService{repo: repo}
}

func (s *TaskListService) Create(userId int, list rsapi.TaskList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TaskListService) GetAll(userId int) ([]rsapi.TaskList, error) {
	return s.repo.GetAll(userId)
}

func (s *TaskListService) GetById(userId, listId int) (rsapi.TaskList, error) {
	return s.repo.GetById(userId, listId)
}
