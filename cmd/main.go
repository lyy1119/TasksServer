package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	myapi "github.com/lyy1119/TasksServer/internal/api"
	"github.com/lyy1119/TasksServer/internal/config"
	"github.com/lyy1119/TasksServer/internal/db"
	"github.com/lyy1119/TasksServer/internal/openapi"
)

func main() {
	portAssign := ":8080"

	if len(os.Args) >= 3 && (os.Args[1] == "-p" || os.Args[1] == "-P") {
		// main -p 8080
		// 尝试将第三个参数转化为整数类型
		port, err := strconv.Atoi(os.Args[2])
		if err != nil {
			port = 8080
			fmt.Printf("Invalid Port \"%s\", use %s as default\n", os.Args[2], portAssign)
		}
		portAssign = fmt.Sprintf(":%d", port)
	}

	r := chi.NewRouter()
	server := myapi.NewServer()
	// h := openapi.HandlerFromMux(server, r)

	r.Route("/api/v1", func(sub chi.Router) {
		// HandlerFromMux
		h := openapi.HandlerFromMux(server, sub)
		sub.Mount("/", h)
	})

	fmt.Println("Loading Config...")
	cfg := config.GetConfig()

	fmt.Println("Connecting MYSQL...")

	db.Open(context.Background(), db.Config{
		DSN:             db.BuildDSN(cfg.MysqlAccount, cfg.MysqlPassword, cfg.MysqlAddress, cfg.MysqlDBName),
		MaxOpenConns:    20,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		PingTimeout:     3 * time.Second})

	fmt.Println("Server started at", portAssign)
	log.Fatal(http.ListenAndServe(portAssign, r))

}
