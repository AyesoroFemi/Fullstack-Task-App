package main

import (
	"errors"
	"fmt"
	"net/http"
	"task-app/internal/repository"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type application struct {
	config config
	store repository.TaskStorage
}

type config struct {
	maxOpenConns int
	maxIdleConns int
	maxIdleTime string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	
	r.Route("/health", func(r chi.Router) {
		r.Get("/", app.healthCheckHandler)
	})
	r.Route("/v1/tasks", func(r chi.Router) {
		r.Get("/", app.getTaskHandler)
		r.Post("/", app.CreateTaskHandler)
		r.Put("/{id}", app.toggleTaskHandler)
		r.Delete("/{id}", app.deleteHandler)
	})

	return r
}


func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	fmt.Println("Served has started", "addr", ":8080")	

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
