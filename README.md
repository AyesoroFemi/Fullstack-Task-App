
# 📝 Task App - Golang (Chi Router) + React (Vite)

A simple full-stack Task Management Application built with **Go** (using Chi Router, SQLite, and clean architecture principles) and **React JS** (Vite) for the frontend.

The purpose of this project is to learn and demonstrate concepts like:
- Clean folder structure
- Dependency Injection
- Interfaces and repository patterns
- RESTful API development
- Simple frontend integration using React

---

## 📂 Project Folder Structure

```bash
/task-app-backend
│
├── cmd/                          # Entry point
│   ├── main.go                   # Application startup
│   ├── handler.go                # Handler functions called from route configuration
│   └── api.go                    # Route configuration and middleware setup
│
├── internal/
│   ├── repository/              # SQL functions implementation
│   │   ├── dbrepo.go            # SQL queries and logic implementation
│   │   └── repository.go        # Repository interface definitions
│   ├── db/                      # Database logic and connection helpers
│   │   └── db.go                # Database initialization and connection setup
│   └── model/                   # Struct definitions
│       └── model.go             # Data models (e.g., Task struct)
│
├── web/                         # Frontend React JS (Vite)
│   ├── index.html
│   └── src/
│       ├── main.jsx             # Entry point for React app
│       └── components/          # Reusable UI components
│
├── go.mod
└── go.sum
```

---

## 📌 REST API Routes

In your `cmd/api.go`, routes are defined using **Chi Router** like so:

```go
r.Route("/v1/tasks", func(r chi.Router) {
    r.Get("/", app.getTaskHandler)         // Get all tasks
    r.Post("/", app.CreateTaskHandler)     // Create new task
    r.Put("/{id}", app.toggleTaskHandler)  // Toggle task completion
    r.Delete("/{id}", app.deleteHandler)   // Delete a task
})
```

---

## ▶️ Running the Backend (Go API)

```bash
cd task-app
go run ./cmd
```

---

## ⚛️ Running the Frontend (React + Vite)

```bash
cd web
npm install      # Install node modules
npm run dev      # Start development server
```

---

## 📖 Learning Outcomes

This project helps reinforce:
- How to structure a Go project cleanly
- Abstracting logic using interfaces
- Writing maintainable code
- Building a full-stack application from backend to frontend

---
