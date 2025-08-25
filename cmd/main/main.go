package main

import (
	"context"
	"fmt"

	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	myapi "github.com/lyy1119/TasksServer/internal/api"
	"github.com/lyy1119/TasksServer/internal/config"
	"github.com/lyy1119/TasksServer/internal/db"
	"github.com/lyy1119/TasksServer/internal/openapi"
	log "github.com/sirupsen/logrus"
)

func main() {
	// log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.Info("Starting Server...")

	log.Info("Loading config from ENV...")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Warn(fmt.Sprintf("Loading config error, maybe wrong formate env variable: %s", err))
	}
	log.Info("Loading config complete.")

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Info("Server will start at", port)

	database, err := db.Open(context.Background(), db.Config{
		DSN:             db.BuildDSN(cfg.MysqlAccount, cfg.MysqlPassword, cfg.MysqlAddress, cfg.MysqlDBName),
		MaxOpenConns:    20,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		PingTimeout:     3 * time.Second,
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("MySQL connecting Error: %s", err))
	}
	log.Info("Database connecting success.")

	log.Info("Starting Router...")
	r := chi.NewRouter()
	server := myapi.NewServer(database)

	r.Route("/api/v1", func(sub chi.Router) {
		// HandlerFromMux
		h := openapi.HandlerFromMux(server, sub)
		sub.Mount("/", h)
	})

	log.Info("Server started.")
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}

}
