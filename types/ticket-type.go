package types

type TicketType struct {
	ID    string  `db:"id"    json:"id"`
	Name  string  `db:"name"  json:"name"`
	Price float64 `db:"price" json:"price"`
	Color string  `db:"color" json:"color"`
}
