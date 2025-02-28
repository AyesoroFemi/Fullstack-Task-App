package main

import (
	"fmt"
	"log"
	"task-app/internal/db"
	"task-app/internal/repository"
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

	store := repository.NewStorage(db)

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