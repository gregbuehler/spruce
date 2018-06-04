package accounts

import (
		// database
		"database/sql"
)

type Account struct {
	accountIdentifier string `json:"accountIdentifier, required" description:"Account Identifier"`
	accountNumber  string `json:"accountNumber , required" description:"Account Number"`
}

func AllAccounts(db *sql.DB) ([]*Account, error) {
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accts := make([]*Account, 0)
	for rows.Next() {
		acct := new(Account)
		err := rows.Scan(
			&acct.accountIdentifier, 
			&acct.accountNumber,
		)
		if err != nil {
			return nil, err
		}
		accts = append(accts, acct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accts, nil
}
