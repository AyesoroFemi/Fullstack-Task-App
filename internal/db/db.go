package db

import (
	// "context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func New( maxOpenConns int, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "api.db") 
	if err != nil {
		return nil, err
	}

	if maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	}

	db.SetMaxIdleConns(maxIdleConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	createTable(db)

	return db, nil
}


func createTable (db *sql.DB) {
	createTaskTable := `
	CREATE TABLE IF NOT EXISTS task_lists (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
	    completed INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME	DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(createTaskTable)

	if err != nil {
		log.Fatal(err)
	}

}

