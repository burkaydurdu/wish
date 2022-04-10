package wish

type Repository interface {
	CreateWish(wish *CreateWishRequest) error
	GetWishByEmail(email string) (*Wish, error)
}
