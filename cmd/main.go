package main

import (
	"task-app/internal/db"
	"task-app/internal/repo"
	"fmt"
	"log"
)


func main() {

	dbConfig := config {
		maxOpenConns: 30,
		maxIdleConns: 30,
		maxIdleTime: "15m",
	}

	// Main database
	db, err := db.New(
		dbConfig.maxOpenConns,
		dbConfig.maxIdleConns,
		dbConfig.maxIdleTime,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := repo.NewStorage(db)

	app := &application{
		config:  dbConfig,
		store: store,
	}

	mux := app.mount()
    if err := app.run(mux); err != nil {
		fmt.Println("err connecting ")
    }
	log.Println(app.run(mux))

}