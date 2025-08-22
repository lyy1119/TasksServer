package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://127.0.0.1:8080/healthz")
	if err != nil || resp.StatusCode != http.StatusOK {
		os.Exit(1) // unhealthy
	}
	os.Exit(0) // healthy
}
