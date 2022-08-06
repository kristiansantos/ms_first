package mongodb

import (
	"context"
	"fmt"
	"html"
	"sync"

	"github.com/kristiansantos/ms_first/pkg/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	singletonStorage *Storage = nil
	once             sync.Once
)

type Storage struct {
	Error   error
	MongoDB *mongo.Database
}

func New(ctx context.Context) Storage {
	environment := *env.SingletonInstance

	if singletonStorage == nil {
		mongoDB, err := connect(ctx, environment)

		once.Do(
			func() {
				singletonStorage = &Storage{
					Error:   err,
					MongoDB: mongoDB,
				}
			})
	}

	return *singletonStorage
}

func connect(ctx context.Context, init env.Application) (*mongo.Database, error) {
	connString := fmt.Sprintf("mongodb://%s:%s@%s/%s", init.Mongo.User, init.Mongo.Pass, init.Mongo.Host, init.Mongo.Database)

	if init.Mongo.Args != "" {
		connString = fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", init.Mongo.User, init.Mongo.Pass, init.Mongo.Host, init.Mongo.Database, init.Mongo.Args)
	}

	mongodbUri := html.UnescapeString(connString)

	clientOptions := options.Client().ApplyURI(mongodbUri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(init.Mongo.Database), nil
}
