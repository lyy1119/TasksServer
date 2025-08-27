// internal/repository/task.go
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type Task struct {
	ID          int64
	Title       string
	Description *string
	CreatedAt   time.Time
}

func ListTasks(ctx context.Context, db *sql.DB, limit int) ([]Task, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := db.QueryContext(ctx,
		`SELECT id, title, description, created_at
		   FROM tasks
		  ORDER BY id DESC
		  LIMIT ?`, limit)
	if err != nil {
		log.Warn(fmt.Sprintf("A error occrus when query in database. Detail: \"%s\"", err))
		return nil, err
	}
	defer rows.Close()

	var out []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func GetTaskByID(ctx context.Context, db *sql.DB, id int64) (*Task, error) {
	var t Task
	err := db.QueryRowContext(ctx,
		`SELECT id, title, description, created_at
		   FROM tasks
		  WHERE id = ?`, id).
		Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func CreateNewTask(ctx context.Context, db *sql.DB, task Task) {

}
