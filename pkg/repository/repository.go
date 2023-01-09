package repository

import (
	"github.com/jmoiron/sqlx"
	rsapi "github.com/wanna-beat-by-bit/rest-golang"
)

type Authorization interface {
	CreateUser(user rsapi.User) (int, error)
}

type TaskList interface {
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
	}
}
