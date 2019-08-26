package api

import (
	"time"
	"fmt"
	"errors"
)

type apiResponse uint
const ResponseHTML apiResponse = 0
const ResponseJSON apiResponse = 1

type TransactionType string

const TransactionTypeDebit TransactionType = "debit"
const TransactionTypeCredit TransactionType = "credit"

func (tt TransactionType) validate() error {

	switch tt {
	case TransactionTypeDebit: return nil
	case TransactionTypeCredit: return nil
	}
	return errors.New(fmt.Sprintf("unkown type '%s", tt))
}

type Money float64

func (m Money) validate() error {

	if m != 0 {
		return nil
	}
	return errors.New(fmt.Sprintf("invalid amount '%f", m))
}

type Transaction struct {
	Type   TransactionType `json:"type"`
	Amount Money           `json:"amount"`
}

type TransactionRecord struct {
	Transaction
	ID     string          `json:"id"`
	Date   time.Time       `json:"effectiveDate"`
}

type TransactionResponse struct {
	ID     string          `json:"id"`
}

type Account struct {
	CurrentBalance  Money    `json:"currentAccountBalance"`

}


type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}


