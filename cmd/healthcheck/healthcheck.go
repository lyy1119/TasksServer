package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{Timeout: 2 * time.Second}
	port := os.Getenv("PORT")
	resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:%s/api/v1/healthz", port))
	if err != nil || resp.StatusCode != http.StatusOK {
		os.Exit(1) // unhealthy
	}
	os.Exit(0) // healthy
}
