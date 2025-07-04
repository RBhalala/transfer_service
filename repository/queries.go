package repository

// Account 

const (
	QCreateAccount        = `INSERT INTO accounts (account_id, balance) VALUES ($1, $2)`
	QGetAccount           = `SELECT balance FROM accounts WHERE account_id = $1`
	QGetAccountForUpdate  = `SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE`
	QUpdateAccountBalance = `UPDATE accounts SET balance = $1 WHERE account_id = $2`
)

// Transaction 

const (
	QInsertTransaction = `
		INSERT INTO transactions (source_account_id, destination_account_id, amount)
		VALUES ($1, $2, $3)`
)
