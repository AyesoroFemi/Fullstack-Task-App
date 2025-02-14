package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"task-app/internal/model"
	"github.com/go-chi/chi/v5"
)


func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) 

	data := map[string]string{
		"status":  "ok",
		"env":     "8080",
		"version": "1.1.0",
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) toggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = app.store.TaskStorage.ToggleTaskCompletion(id)
	if err != nil {
		log.Println("Error toggling task completion:", err)
		http.Error(w, "Failed to toggle task", http.StatusInternalServerError)
		return
	}

	toggleTask, err := app.store.TaskStorage.GetSingleTask(id)
	if err != nil {
		log.Println("Error fetching updated task:", err)
		http.Error(w, "Failed to fetch updated task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(toggleTask); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}


func (app *application) getTaskHandler (w http.ResponseWriter, r *http.Request) {
	getTasks, err := app.store.TaskStorage.GetAllTasks()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(getTasks)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
}


func (app *application) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t struct {
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println(err)
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	task := &model.Task{
		Title:     t.Title,
		Completed: t.Completed == 1,
	}

	err = app.store.TaskStorage.CreateTask(task)
	if err != nil {
		log.Println(err)
		http.Error(w, `{"error": "Failed to create task"}`, http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
}



func (app *application) deleteHandler (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = app.store.TaskStorage.DeleteTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}