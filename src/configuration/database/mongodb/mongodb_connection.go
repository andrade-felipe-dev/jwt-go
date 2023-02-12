package mongodb

import (
	"context"
	"os"

	"github.com/andrade-felipe-dev/jwt-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL  = "MONGODB_URL"
	MONGODB_USER = "MONGODB_USER"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri),
	)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Conexão realizada com sucesso no banco de dados!")

	return client.Database(mongodb_database), nil
}
