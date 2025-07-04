package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"transfer_service/service"
)

type Handler struct {
	transferService service.ITransferService
}

func TransferHandler(s service.ITransferService) *Handler {
	return &Handler{transferService: s}
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID      int64  `json:"account_id"`
		InitialBalance string `json:"initial_balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.transferService.CreateAccount(r.Context(), req.AccountID, req.InitialBalance)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "account_id")
	accountID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid account id")
		// http.Error(w, , http.StatusBadRequest)
		return
	}

	balance, err := h.transferService.GetAccount(r.Context(), accountID)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, err.Error())
		return
	}

	resp := struct {
		AccountID int64  `json:"account_id"`
		Balance   string `json:"balance"`
	}{
		AccountID: accountID,
		Balance:   balance.Text('f', 10),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) SubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SourceAccountID      int64  `json:"source_account_id"`
		DestinationAccountID int64  `json:"destination_account_id"`
		Amount               string `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.transferService.SubmitTransaction(r.Context(), req.SourceAccountID, req.DestinationAccountID, req.Amount)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
