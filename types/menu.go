package types

type Menu struct {
	Article          `db:""`
	IncludedTickets  []string ` json:"includedTickets" db:"tickets"`
	IncludedArticles []string `json:"includedArticles" db:"items"`
}
