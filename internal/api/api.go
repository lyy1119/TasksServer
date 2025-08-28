package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lyy1119/TasksServer/internal/repository"
)

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
func httpError(w http.ResponseWriter, err error, code int) {
	writeJSON(w, code, map[string]any{"error": err.Error()})
}

type Server struct {
	DB *sql.DB
} // application interface

// data struct
type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewServer(db *sql.DB) Server {
	return Server{
		DB: db,
	}
}

// Health check
// (GET /healthz)
func (Server) GetHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ok":true}`))
}

// Get all tasks
// (GET /tasks)
func (s Server) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second) // 最多3秒查询操作
	defer cancel()

	// tasks is a slice
	tasks, err := repository.ListTasks(ctx, s.DB, 100)
	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

// Create a new task
// (POST /tasks)
func (s Server) PostTasks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var in repository.Task
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	repository.CreateNewTask(ctx, s.DB, in)
	writeJSON(w, http.StatusCreated, map[string]bool{"ok": true})
}

// Delete a task
// (DELETE /tasks/{id})
func (Server) DeleteTasksId(w http.ResponseWriter, r *http.Request, id int) {

}

// Get a task by ID
// (GET /tasks/{id})
func (Server) GetTasksId(w http.ResponseWriter, r *http.Request, id int) {

}

// Update a task
// (PUT /tasks/{id})
func (Server) PutTasksId(w http.ResponseWriter, r *http.Request, id int) {

}

// Get all users
// (GET /users)
func (Server) GetUsers(w http.ResponseWriter, r *http.Request) {

}

// Create a new user
// (POST /users)
func (Server) PostUsers(w http.ResponseWriter, r *http.Request) {

}

// Delete a user
// (DELETE /users/{id})
func (Server) DeleteUsersId(w http.ResponseWriter, r *http.Request, id int) {

}

// Get a single user by ID
// (GET /users/{id})
func (Server) GetUsersId(w http.ResponseWriter, r *http.Request, id int) {

}

// Update a user
// (PUT /users/{id})
func (Server) PutUsersId(w http.ResponseWriter, r *http.Request, id int) {

}
