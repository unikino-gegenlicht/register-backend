package types

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type DBTransaction struct {
	ID        string             `db:"id"        json:"id"`
	Timestamp pgtype.Timestamptz `db:"timestamp" json:"timestamp"`
	JWT       string             `db:"data"      json:"data"`
}

func (dbT DBTransaction) ExtractTransactionData() (*TransactionData, error) {
	return nil, nil
}

type TransactionData struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	PaymentType string    `json:"paymentType"`
	Amount      float64   `json:"amount"`
	Tickets     []string  `json:"tickets"`
	Articles    []Article `json:"articles"`
	Menus       []Menu    `json:"menu"`
}
