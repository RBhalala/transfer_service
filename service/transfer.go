package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"

	"transfer_service/repository"
)


type transferService struct {
	db   *sql.DB
	repo repository.ITransferRepo
}

func NewTransferService(db *sql.DB) ITransferService {
	return &transferService{db: db, repo: repository.NewPostgresTransferRepo()}
}


func (s *transferService) CreateAccount(ctx context.Context, id int64, initial string) error {
	bal, err := repository.ParseNonNegativeDecimal(initial)
	if err != nil {
		return err
	}
	return s.repo.CreateAccount(ctx, s.db, id, bal)
}

func (s *transferService) GetAccount(ctx context.Context, id int64) (*big.Float, error) {
	return s.repo.GetAccount(ctx, s.db, id)
}

func (s *transferService) SubmitTransaction(ctx context.Context, srcID, dstID int64, amtStr string) (err error) {
	amt, err := repository.ParsePositiveDecimal(amtStr)
	if err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	if srcID == dstID {
		return errors.New("Source and destination accounts must be different to transfer funds.")
	}
	srcBal, err := s.repo.GetAccountForUpdate(ctx, tx, srcID)
	if err != nil {
		return errors.New("Source account not found")
	}
	dstBal, err := s.repo.GetAccountForUpdate(ctx, tx, dstID)
	if err != nil {
		return errors.New("Destination account not found")
	}

	if srcBal.Cmp(amt) < 0 {
		return errors.New("insufficient funds")
	}

	newSrc := new(big.Float).Sub(srcBal, amt)
	newDst := new(big.Float).Add(dstBal, amt)

	if err = s.repo.UpdateBalance(ctx, tx, srcID, newSrc); err != nil {
		return
	}
	if err = s.repo.UpdateBalance(ctx, tx, dstID, newDst); err != nil {
		return
	}
	err = s.repo.InsertTransaction(ctx, tx, srcID, dstID, amt)
	return
}
