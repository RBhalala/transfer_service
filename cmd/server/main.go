package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"transfer_service/service"
	"transfer_service/handler"
)

func main() {
	// Replace with your Postgres connection string
	dsn := "postgres://root:root@localhost:5432/transfer_service?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	transferSvc := service.NewTransferService(db)
	handler := handler.NewHandler(transferSvc)

	r := chi.NewRouter()

	r.Post("/accounts", handler.CreateAccount)
	r.Get("/accounts/{account_id}", handler.GetAccount)
	r.Post("/transactions", handler.SubmitTransaction)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
