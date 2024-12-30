package types

type Menu struct {
	Article          `db:""`
	IncludedTickets  []string `db:"tickets" json:"includedTickets"`
	IncludedArticles []string `db:"items"   json:"includedArticles"`
}
