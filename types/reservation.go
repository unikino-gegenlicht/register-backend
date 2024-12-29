package types

import "time"

type Reservation struct {
	ID                string    `json:"id" db:"id"`
	CreatedAt         time.Time `json:"createdAt" db:"created_at"`
	PickedUpAt        time.Time `json:"pickedUpAt" db:"picked_up_at"`
	HolderInformation struct {
		FirstName string `json:"firstName" db:"first_name"`
		LastName  string `json:"lastName" db:"last_name"`
		EMail     string `json:"email" db:"email"`
	}
	Screening string `json:"screening" db:"screening"`
	Seats     int    `json:"seats" db:"seats"`
}
