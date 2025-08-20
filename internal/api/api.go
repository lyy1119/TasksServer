package api

import (
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// Health check
// (GET /healthz)
func (Server) GetHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ok":true}`))
}

// Get all tasks
// (GET /tasks)
func (Server) GetTasks(w http.ResponseWriter, r *http.Request) {

}

// Create a new task
// (POST /tasks)
func (Server) PostTasks(w http.ResponseWriter, r *http.Request) {

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
