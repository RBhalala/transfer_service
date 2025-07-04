package server

import (
	"log"
	"net/http"
	"database/sql"

	"transfer_service/handler"
	"transfer_service/service"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func RegisterRouters(db *sql.DB) http.Handler {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: no .env file found")
	}

	transferService := service.TransferService(db)
	h := handler.NewHandler(transferService)

	r := chi.NewRouter()
	r.Post("/accounts", h.CreateAccount)
	r.Get("/accounts/{account_id}", h.GetAccount)
	r.Post("/transactions", h.SubmitTransaction)

	return r
}
