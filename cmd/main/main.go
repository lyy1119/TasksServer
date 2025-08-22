package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	myapi "github.com/lyy1119/TasksServer/internal/api"
	"github.com/lyy1119/TasksServer/internal/config"
	"github.com/lyy1119/TasksServer/internal/db"
	"github.com/lyy1119/TasksServer/internal/openapi"
)

func main() {

	fmt.Println("Loading Config...")
	cfg := config.GetConfig()

	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Println("Notice: Server will start at", port)

	r := chi.NewRouter()
	server := myapi.NewServer()
	// h := openapi.HandlerFromMux(server, r)

	r.Route("/api/v1", func(sub chi.Router) {
		// HandlerFromMux
		h := openapi.HandlerFromMux(server, sub)
		sub.Mount("/", h)
	})

	fmt.Println("Connecting MYSQL...")

	db.Open(context.Background(), db.Config{
		DSN:             db.BuildDSN(cfg.MysqlAccount, cfg.MysqlPassword, cfg.MysqlAddress, cfg.MysqlDBName),
		MaxOpenConns:    20,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		PingTimeout:     3 * time.Second})

	log.Fatal(http.ListenAndServe(port, r))

}
