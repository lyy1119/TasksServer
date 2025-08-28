// internal/repository/task.go
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Task struct {
	ID          int64
	Title       string
	Description *string
	CreatedAt   time.Time
}

// a == b return true
func compareTasks(a *Task, b *Task) bool {
	return (a.Title == b.Title && a.Description == b.Description)
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
		log.Warn(fmt.Sprintf("A error occrus when query in database. Detail: \"%s\"", err))
		return nil, nil
	}
	if err != nil {
		log.Warn(fmt.Sprintf("A error occrus when query in database. Detail: \"%s\"", err))
		return nil, err
	}
	return &t, nil
}

func CreateNewTask(ctx context.Context, db *sql.DB, task Task) error {
	if strings.TrimSpace(task.Title) == "" {
		log.Warn("No title when create a new task.")
		return errors.New("title is required")
	}
	// 执行insert语句
	res, err := db.ExecContext(ctx,
		`INSERT INTO tasks (title, description) VALUES (?, ?)`,
		task.Title, task.Description)
	if err != nil {
		log.Warn(fmt.Sprintf("A error occur when create a new task. Detial: %s", err))
		return err
	}
	// 获取插入的id
	id, _ := res.LastInsertId()
	// 获取插入的任务来验证
	newtask, err := GetTaskByID(ctx, db, id)
	if err != nil {
		log.Warn(fmt.Sprintf("A error occur when query after a INSERT. Detail: %s", err))
		return err
	}
	// 验证插入的数据
	if !compareTasks(&task, newtask) {
		log.Warn("Not the same, some problems may be occured between backend and Mysql")
		DeleteTasksByID(ctx, db, id) // ?
		return errors.New("insert failed, internal error")
	}

	return nil
}

func DeleteTasksByID(ctx context.Context, db *sql.DB, id int64) (bool, error) {
	return true, nil
}
