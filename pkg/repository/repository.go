package repository

import (
	"github.com/jmoiron/sqlx"
	rsapi "github.com/wanna-beat-by-bit/rest-golang"
)

type Authorization interface {
	CreateUser(user rsapi.User) (int, error)
	GetUser(username, password string) (rsapi.User, error)
}

type TaskList interface {
	Create(userId int, list rsapi.TaskList) (int, error)
	GetAll(userId int) ([]rsapi.TaskList, error)
	GetById(userId, listId int) (rsapi.TaskList, error)
}

type TaskItem interface {
}

type Repository struct {
	Authorization
	TaskItem
	TaskList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TaskList:      NewTaskListPostgres(db),
	}
}
