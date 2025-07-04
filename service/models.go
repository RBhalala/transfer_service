package service

import "math/big"

type Account struct {
	AccountID int64
	Balance   *big.Float
}

type Transaction struct {
	SourceAccountID      int64
	DestinationAccountID int64
	Amount               *big.Float
}
