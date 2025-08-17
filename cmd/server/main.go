package main

import (
	"log"
	"schoolapi/internal/api/listeners"
)

func main() {
	err := listeners.StartAPIServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}