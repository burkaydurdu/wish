package couchbase

import (
	"fmt"

	"github.com/burkaydurdu/wish/internal/domain/wish"
	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

func CreateWishRepositoryConfig(client *Client, bucket *gocb.Bucket, scope *gocb.Scope, logger *zap.Logger) *NewCouchbaseRepositoryOpts {
	return &NewCouchbaseRepositoryOpts{
		Client:     client,
		Bucket:     bucket,
		Scope:      scope,
		Collection: scope.Collection(wishCollectionName),
		Logger:     logger,
	}
}

func NewWishRepository(opts *NewCouchbaseRepositoryOpts) wish.Repository {
	return &repository{
		logger:     opts.Logger,
		collection: opts.Collection,
		scope:      opts.Scope,
		client:     opts.Client,
	}
}

func (r *repository) CreateWish(wish *wish.CreateWishRequest) error {
	_, err := r.collection.Insert(wish.ID, wish, nil)

	return err
}

func (r *repository) GetWishByEmail(email string) (*wish.Wish, error) {
	query := fmt.Sprintf("SELECT %s.* FROM %s WHERE email = '%s'", wishCollectionName, wishCollectionName, email)
	rows, err := r.scope.Query(query, nil)

	if err != nil {
		return nil, err
	}

	var wishData *wish.Wish

	for rows.Next() {
		err = rows.Row(&wishData)

		if err != nil {
			return nil, err
		}

		break
	}

	return wishData, nil
}
