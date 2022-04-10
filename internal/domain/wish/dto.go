package wish

import (
	"time"
)

type CreateWishRequest struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Wish      string    `json:"wish"`
	CreatedAt time.Time `json:"created_at"`
}

type Error struct {
	Message string
}

type CreateWishResponse struct {
	Message string `json:"message"`
}
