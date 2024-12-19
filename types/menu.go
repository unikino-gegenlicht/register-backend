package types

type Menu struct {
	Article          `db:""`
	IncludedArticles []string `json:"includedArticles" db:"included_articles"`
}
