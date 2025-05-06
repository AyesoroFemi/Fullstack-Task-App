package repository

import (
	"database/sql"
	"task-app/internal/model"
)


// type TaskRepo struct {
// 	TaskStorage interface {
// 		GetAllTasks() ([]model.Task, error)
// 		CreateTask(task *model.Task) error
// 		ToggleTaskCompletion(taskID int) error
// 		DeleteTask(id int) error
// 		GetSingleTask(id int) (*model.Task, error)
// 	}
// }


// func NewStorage(db *sql.DB) TaskRepo {
// 	return TaskRepo{
// 		TaskStorage: &TaskStore{db},
// 	}
// }


type TaskStorage interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task *model.Task) error
	ToggleTaskCompletion(taskID int) error
	DeleteTask(id int) error
	GetSingleTask(id int) (*model.Task, error)
}

func NewStorage(db *sql.DB) TaskStorage {
	return &TaskStore{ db: db}
}