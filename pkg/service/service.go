package service

import (
	rsapi "github.com/wanna-beat-by-bit/rest-golang"
	"github.com/wanna-beat-by-bit/rest-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user rsapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TaskList interface {
	Create(userId int, list rsapi.TaskList) (int, error)
	GetAll(userId int) ([]rsapi.TaskList, error)
	GetById(userId, listId int) (rsapi.TaskList, error)
}

type TaskItem interface {
}

type Service struct {
	Authorization
	TaskItem
	TaskList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TaskList:      NewTaskListService(repos.TaskList),
	}
}
