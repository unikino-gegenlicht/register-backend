package types

type Ticket struct {
	ID    string  `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
	Color string  `json:"color" db:"color"`
}
