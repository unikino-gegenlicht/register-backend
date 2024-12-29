package types

type Screening struct {
	ID             string `json:"id" db:"id"`
	WordpressID    int    `json:"wpID" db:"wordpress_id"`
	Title          string `json:"title" db:"title"`
	AvailableSeats int    `json:"availableSeats"`
}
