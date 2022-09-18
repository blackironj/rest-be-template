package repository

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCli *mongo.Client

func OpenMongoDB(url, dbName string) *mongo.Database {
	clientOpts := options.Client().
		ApplyURI(url).
		SetMaxPoolSize(30)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to connect mongodb")
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal().Err(err).Msg("fail to ping to mongodb")
	}

	mongoCli = client

	return client.Database(dbName)
}

func CloseMongoDB() {
	_ = mongoCli.Disconnect(context.Background())
}

func MongoTransaction(txTask func(sc mongo.SessionContext) error) error {
	ctx := context.Background()

	session, err := mongoCli.StartSession()
	if err != nil {
		log.Error().Err(err).Msg("fail to start mongo session")
		return err
	}
	defer session.EndSession(ctx)

	if err := session.StartTransaction(); err != nil {
		log.Error().Err(err).Msg("fail to start mongo tx")
		return err
	}

	if err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err := txTask(sc); err != nil {
			if abortErr := session.AbortTransaction(sc); err != nil {
				log.Error().Err(err).Msg("mongo tx abort err")
				return abortErr
			}
			return err
		}

		if commitErr := session.CommitTransaction(sc); commitErr != nil {
			log.Error().Err(err).Msg("mongo tx commit err")
			return commitErr
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
