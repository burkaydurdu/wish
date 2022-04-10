package postgres

import (
	"fmt"

	"github.com/burkaydurdu/wish/internal/domain/wish"
	"go.uber.org/zap"
)

func NewWishRepository(client *Client, logger *zap.Logger) wish.Repository {
	return &repository{
		logger: logger,
		client: client,
	}
}

func (r *repository) CreateWish(wish *wish.CreateWishRequest) error {
	query := fmt.Sprintf("INSERT into %s(email, wish, created_at) VALUES ($1, $2, $3)", wishTableName)
	_, err := r.client.DB.Exec(query, wish.Email, wish.Wish, wish.CreatedAt)
	return err
}

func (r *repository) GetWishByEmail(_ string) (*wish.Wish, error) {
	return nil, nil
}
