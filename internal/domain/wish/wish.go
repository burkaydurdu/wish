package wish

import "time"

type Wish struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Wish      string    `json:"wish"`
	CreatedAt time.Time `json:"created_at"`
}
