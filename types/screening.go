package types

type Screening struct {
	ID             string `db:"id"               json:"id"`
	WordpressID    int    `db:"wordpress_id"     json:"wpID"`
	Title          string `db:"title"            json:"title"`
	AvailableSeats int    `json:"availableSeats"`
}
