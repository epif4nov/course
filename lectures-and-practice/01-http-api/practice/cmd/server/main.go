package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"http-api-practice/internal/app"
)

func main() {
	address := os.Getenv("HTTP_ADDR")
	if address == "" {
		address = ":8080"
	}

	server := &http.Server{
		Addr:              address,
		Handler:           app.New(log.New(os.Stdout, "http ", log.LstdFlags)),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("HTTP API listening on %s", address)
	log.Fatal(server.ListenAndServe())
}
