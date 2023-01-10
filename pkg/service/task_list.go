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
