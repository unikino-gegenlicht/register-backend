package types

import "github.com/jackc/pgx/v5/pgtype"

type Transaction struct {
	ID        string             `db:"id"        json:"id"`
	Timestamp pgtype.Timestamptz `db:"timestamp" json:"timestamp"`
	Amount    float64            `db:"amount"    json:"amount"`
	Tickets   []string           `db:"tickets"   json:"tickets"`
	Menus     []string           `db:"menus"     json:"menus"`
	Articles  []string           `db:"articles"  json:"articles"`
}
