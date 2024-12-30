package types

import "time"

type Reservation struct {
	ID                string    `db:"id"           json:"id"`
	CreatedAt         time.Time `db:"created_at"   json:"createdAt"`
	PickedUpAt        time.Time `db:"picked_up_at" json:"pickedUpAt"`
	HolderInformation struct {
		FirstName string `db:"first_name" json:"firstName"`
		LastName  string `db:"last_name"  json:"lastName"`
		EMail     string `db:"email"      json:"email"`
	}
	Screening string `db:"screening" json:"screening"`
	Seats     int    `db:"seats"     json:"seats"`
}
