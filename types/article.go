package types

type Article struct {
	ID      string `binding:"optional" db:"id"      json:"id"`
	Name    string `binding:"required" db:"name"    json:"name"`
	Enabled bool   `binding:"enabled"  db:"enabled" json:"enabled"`
	Prices  struct {
		Members *float64 `binding:"required" db:"price_members" json:"members"`
		Guests  float64  `binding:"required" db:"price_guests"  json:"guests"`
	} `json:"prices" db:""`
	Color string `binding:"iscolor,default=#00a2ff" db:"color" json:"color"`
}
