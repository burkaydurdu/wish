package wish

import (
	"time"

	"github.com/burkaydurdu/wish/internal/validations"
)

type Service interface {
	CreateWish(*CreateWishRequest) error
	GetWishByEmail(email string) (*Wish, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) CreateWish(requestDTO *CreateWishRequest) error {
	//wishData, err := s.GetWishByEmail(requestDTO.Email)

	//if err != nil {
	//	return err
	//}

	//if wishData != nil {
	//	return ErrEmailAlreadyExists
	//}

	//id, err := uuid.NewUUID()

	//if err != nil {
	//	return err
	//}

	// Set id and current time for a wish
	//requestDTO.ID = id.String()
	requestDTO.CreatedAt = time.Now()

	err := s.r.CreateWish(requestDTO)

	return err
}

func (s *service) GetWishByEmail(email string) (*Wish, error) {
	if !validations.IsEmail(email) {
		return nil, ErrInvalidEmail
	}

	result, err := s.r.GetWishByEmail(email)

	if err != nil {
		return nil, err
	}

	return result, nil
}
