package types

type Article struct {
	ID      string `db:"id"            json:"id"`
	Name    string `binding:"required" db:"name"      json:"name"`
	Enabled bool   `db:"enabled"       json:"enabled"`
	Prices  struct {
		Members *float64 `db:"price_members" json:"members"`
		Guests  float64  `binding:"required" db:"price_guests" json:"guests"`
	} `json:"prices" db:""`
	Color string `binding:"omitempty,hexcolor|rgb|hsl" db:"color" json:"color"`
}
