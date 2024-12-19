package types

import "github.com/jackc/pgx/v5/pgtype"

type Transaction struct {
	ID        string             `json:"id" db:"id"`
	Timestamp pgtype.Timestamptz `json:"timestamp" db:"timestamp"`
	Amount    float64            `json:"amount" db:"amount"`
	Tickets   struct {
		Standard uint `json:"standard" db:"standard_tickets"`
		Free     uint `json:"free" db:"free_tickets"`
		Combined uint `json:"combined" db:"combined_tickets"`
	}
	Menus    []string `json:"menus" db:"menus"`
	Articles []string `json:"articles" db:"articles"`
}
