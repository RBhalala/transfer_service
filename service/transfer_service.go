package service

import (
	"context"
	"math/big"
)

type ITransferService interface {
	CreateAccount(ctx context.Context, id int64, initial string) error
	GetAccount(ctx context.Context, id int64) (*big.Float, error)
	SubmitTransaction(ctx context.Context, srcID, dstID int64, amount string) error
}