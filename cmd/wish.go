package cmd

import (
	"context"

	"github.com/burkaydurdu/wish/internal/persistence/postgres"

	"github.com/burkaydurdu/wish/config"
	"github.com/burkaydurdu/wish/internal/domain/wish"
	"go.uber.org/fx"
)

func Run() {
	ctx := context.Background()

	app := fx.New(
		fx.Provide(
			config.New,
			config.GetDatabaseConfig,
			createServer,
			createLogger,
			//couchbase.NewCouchbaseClient,
			//couchbase.CreateBucket,
			//couchbase.CreateScope,
			//couchbase.CreateWishRepositoryConfig,
			//couchbase.NewWishRepository,
			postgres.NewPostgresClient,
			postgres.NewWishRepository,
			wish.NewService),
		fx.Invoke(
			wish.NewHandler,
			config.Print,
			startHttpServer))

	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
