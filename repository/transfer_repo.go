package repository

import (
	"context"
	"math/big"
)

// It Exposes all persistence operations required by the TransferService.
type ITransferRepo interface {
	CreateAccount(ctx context.Context, db IDbtx, id int64, bal *big.Float) error
	GetAccount(ctx context.Context, db IDbtx, id int64) (*big.Float, error)
	GetAccountForUpdate(ctx context.Context, tx IDbtx, id int64) (*big.Float, error)
	UpdateBalance(ctx context.Context, tx IDbtx, id int64, bal *big.Float) error
	InsertTransaction(ctx context.Context, tx IDbtx, src, dst int64, amt *big.Float) error
}
