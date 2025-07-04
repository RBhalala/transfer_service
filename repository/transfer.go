package repository

import (
	"context"
	"math/big"
	"errors"
	"database/sql"
)

type pgTransferRepo struct{}

func NewPostgresTransferRepo() ITransferRepo {
	return &pgTransferRepo{}
}

func (r *pgTransferRepo) CreateAccount(ctx context.Context, db IDbtx, id int64, bal *big.Float) error {
	 // Check if account already exists
	 _, err := r.GetAccount(ctx, db, id)
	 if err == nil {
		 return ErrDupData
	 }
	 if err != ErrNoData {
		 // Unexpected error from GetAccount
		 return err
	 }

	_, err = db.ExecContext(ctx, QCreateAccount, id, bal.Text('f', 10))
	return err
}

func (*pgTransferRepo) GetAccount(ctx context.Context, db IDbtx, id int64) (*big.Float, error) {
	var balStr string
	err := db.QueryRowContext(ctx, QGetAccount, id).Scan(&balStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoData
		}
		return nil, err
	}
	return ParseDecimal(balStr)
}

func (*pgTransferRepo) GetAccountForUpdate(ctx context.Context, tx IDbtx, id int64) (*big.Float, error) {
	var balStr string
	err := tx.QueryRowContext(ctx, QGetAccountForUpdate, id).Scan(&balStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoData
		}
		return nil, err
	}
	return ParseDecimal(balStr)
}

func (*pgTransferRepo) UpdateBalance(ctx context.Context, tx IDbtx, id int64, bal *big.Float) error {
	_, err := tx.ExecContext(ctx, QUpdateAccountBalance, bal.Text('f', 10), id)
	return err
}

func (*pgTransferRepo) InsertTransaction(ctx context.Context, tx IDbtx, src, dst int64, amt *big.Float) error {
	_, err := tx.ExecContext(ctx, QInsertTransaction, src, dst, amt.Text('f', 10))
	return err
}
