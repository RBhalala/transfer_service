package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

// mockTransferService implements service.ITransferService
type mockTransferService struct {
	createAccountFunc    func(ctx context.Context, accountID int64, initialBalance string) error
	getAccountFunc       func(ctx context.Context, accountID int64) (*big.Float, error)
	submitTransactionFunc func(ctx context.Context, sourceAccountID, destinationAccountID int64, amount string) error
}

func (m *mockTransferService) CreateAccount(ctx context.Context, accountID int64, initialBalance string) error {
	return m.createAccountFunc(ctx, accountID, initialBalance)
}

func (m *mockTransferService) GetAccount(ctx context.Context, accountID int64) (*big.Float, error) {
	return m.getAccountFunc(ctx, accountID)
}

func (m *mockTransferService) SubmitTransaction(ctx context.Context, sourceAccountID, destinationAccountID int64, amount string) error {
	return m.submitTransactionFunc(ctx, sourceAccountID, destinationAccountID, amount)
}

//  Global variables


var MockCreateAccount = &mockTransferService{
	createAccountFunc: func(ctx context.Context, accountID int64, initialBalance string) error {
		return nil
	},
}
var TransferCreateAccountHandler = TransferHandler(MockCreateAccount)
var CreateAccountBody = `{"account_id":123,"initial_balance":"100.50"}`
var SubmitTransactionBody = `{"source_account_id":1,"destination_account_id":2,"amount":"50.00"}`
var ErrorRespBody struct {
	Error string `json:"error"`
}

var MockGetAccount = &mockTransferService{
	getAccountFunc: func(ctx context.Context, accountID int64) (*big.Float, error) {
		return big.NewFloat(250.75), nil
	},
}
var TransferGetAccountHandler = TransferHandler(MockGetAccount)

var MockSubmitTransfer = &mockTransferService{
	submitTransactionFunc: func(ctx context.Context, sourceAccountID, destinationAccountID int64, amount string) error {
		return nil
	},
}
var TransferSubmitTransactionHandler = TransferHandler(MockSubmitTransfer)

//  Test functions ------------
// CreateAccount
func TestTransfer_CreateAccount_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBufferString(CreateAccountBody))
	w := httptest.NewRecorder()

	TransferCreateAccountHandler.CreateAccount(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestTransfer_CreateAccount_WithInvalidBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBufferString("invalid json"))
	w := httptest.NewRecorder()

	TransferCreateAccountHandler.CreateAccount(w, req)
	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	
	err := json.NewDecoder(resp.Body).Decode(&ErrorRespBody)
	assert.NoError(t, err)

	expectedErrorMsg := "invalid request body"
	assert.Equal(t, expectedErrorMsg, ErrorRespBody.Error)
}

func TestTransfer_CreateAccount_ValidationFailure(t *testing.T) {
	expectedErrorMsg := "any validation error"

	MockCreateAccount.createAccountFunc = func(ctx context.Context, accountID int64, initialBalance string) error {
		return errors.New(expectedErrorMsg)
	}

	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewBufferString(CreateAccountBody))
	w := httptest.NewRecorder()
	tranHandler := TransferHandler(MockCreateAccount)

	tranHandler.CreateAccount(w, req)
	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	err := json.NewDecoder(resp.Body).Decode(&ErrorRespBody)
	assert.NoError(t, err)

	assert.Equal(t, expectedErrorMsg, ErrorRespBody.Error)
}

// GetAccount
func TestTransfer_GetAccount_success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accounts/123", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("account_id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	TransferGetAccountHandler.GetAccount(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var respBody struct {
		AccountID int64  `json:"account_id"`
		Balance   string `json:"balance"`
	}
	json.NewDecoder(resp.Body).Decode(&respBody)
	assert.Equal(t, int64(123), respBody.AccountID)
	assert.Equal(t, "250.7500000000", respBody.Balance)
}

func TestTransfer_GetAccount_InvalidAccount(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accounts/abc", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("account_id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	TransferGetAccountHandler.GetAccount(w, req)
	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&ErrorRespBody)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "invalid account id", ErrorRespBody.Error)
}

func TestTransfer_GetAccount_AccountNotExist(t *testing.T) {
	MockGetAccount.getAccountFunc = func(ctx context.Context, accountID int64) (*big.Float, error) {
		return nil, errors.New("No Data Found")
	}
	req := httptest.NewRequest(http.MethodGet, "/accounts/123", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("account_id", "123567")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	TransferGetAccountHandler.GetAccount(w, req)
	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&ErrorRespBody)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "No Data Found", ErrorRespBody.Error)
}

// SubmitTransaction 
func TestTransfer_SubmitTransaction_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(SubmitTransactionBody))
	w := httptest.NewRecorder()

	TransferSubmitTransactionHandler.SubmitTransaction(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
func TestTransfer_SubmitTransaction_InvalidJson(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString("invalid json"))
	w := httptest.NewRecorder()

	TransferSubmitTransactionHandler.SubmitTransaction(w, req)
	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&ErrorRespBody)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "invalid request body", ErrorRespBody.Error)
}

func TestTransfer_SubmitTransaction_ValidationFailure(t *testing.T) {
	expectedErrorMsg := "any validation error"
	MockSubmitTransfer.submitTransactionFunc = func(ctx context.Context, sourceAccountID, destinationAccountID int64, amount string) error {
		return errors.New(expectedErrorMsg)
	}
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(SubmitTransactionBody))
	w := httptest.NewRecorder()

	TransferSubmitTransactionHandler.SubmitTransaction(w, req)

	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&ErrorRespBody)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, expectedErrorMsg, ErrorRespBody.Error)
}
