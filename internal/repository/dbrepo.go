package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"task-app/internal/model"
	// "time"
)

// var QueryTimeoutDuration = time.Second * 5

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(t *sql.DB) *TaskStore {
	return &TaskStore{
		db: t,
	}
}


func (t *TaskStore) CreateTask(task *model.Task) error {
	query := `INSERT INTO task_lists (title, completed) VALUES (?, ?)`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}

	completedInt := 0
	if task.Completed {
		completedInt = 1
	}

	_, err = stmt.Exec(task.Title, completedInt)
	if err != nil {
		return err
	}

	return nil

}


func (t *TaskStore) GetAllTasks() ([]model.Task, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM task_lists`
	rows, err := t.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil

}


func (t *TaskStore) GetSingleTask(id int) (*model.Task, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM task_lists WHERE id = ?`

	var task model.Task

	err := t.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task with id %d not found", id)
		}
		return nil, err
	}

	return &task, nil

} 

func (t *TaskStore) ToggleTaskCompletion(taskID int) error {
    query := `UPDATE task_lists SET completed = NOT completed WHERE id = ?`

    res, err := t.db.Exec(query, taskID)
    if err != nil {
        return err
    }

    rows, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if rows == 0 {
        return errors.New("resource not found")
    }

    return nil
}


func (t *TaskStore) DeleteTask(id int) error {
	query := `DELETE FROM task_lists WHERE id = ?`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return nil
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	fmt.Printf("Task with ID %d deleted successfully\n", id)
	
	return nil
}