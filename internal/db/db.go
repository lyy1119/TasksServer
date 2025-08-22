package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DSN             string        // 完整 DSN；或留空用下方 BuildDSN
	MaxOpenConns    int           // 例如 20
	MaxIdleConns    int           // 例如 5
	ConnMaxLifetime time.Duration // 例如 time.Hour
	PingTimeout     time.Duration // 例如 3 * time.Second
}

func BuildDSN(user, pass, addr string, dbname string) string {
	// parseTime/charset/loc 基本是 Web 场景的合理默认
	return user + ":" + pass +
		"@tcp(" + addr + ")/" + dbname +
		"?parseTime=true&charset=utf8mb4&loc=Local"
}

func Open(ctx context.Context, cfg Config) (*sql.DB, error) {
	if cfg.PingTimeout <= 0 {
		cfg.PingTimeout = 3 * time.Second
	}

	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	// 健康检查（带超时）
	pingCtx, cancel := context.WithTimeout(ctx, cfg.PingTimeout)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}
