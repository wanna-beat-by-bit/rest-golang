package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	rsapi "github.com/wanna-beat-by-bit/rest-golang"
)

type TaskListPostgres struct {
	db *sqlx.DB
}

func NewTaskListPostgres(db *sqlx.DB) *TaskListPostgres {
	return &TaskListPostgres{db: db}
}

func (r *TaskListPostgres) Create(userId int, list rsapi.TaskList) (int, error) {
	//Create a user's post in one transaction
	query_manager, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", taskListsTable)
	row := query_manager.QueryRow(createListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		query_manager.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = query_manager.Exec(createUsersListQuery, userId, id)
	if err != nil {
		query_manager.Rollback()
		return 0, err
	}

	return id, query_manager.Commit()
}
