package types

type Article struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
	Prices  struct {
		Members *float64 `json:"members" db:"pice_members"`
		Guests  float64  `json:"guests" db:"price_guests"`
	} `json:"prices" db:""`
	Style struct {
		Color string `json:"color" db:"color"`
		Icon  string `json:"icon" db:"icon"`
	} `json:"style" db:""`
}
