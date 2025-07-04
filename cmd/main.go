package main

import (
	"log"
	"net/http"
	"os"

	"transfer_service/config"
	"transfer_service/server"
)

func main() {
	db, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect DB: %v", err)
    }
    defer db.Close()

	r := server.RegisterRouters(db)
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
	log.Printf("Server started on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
