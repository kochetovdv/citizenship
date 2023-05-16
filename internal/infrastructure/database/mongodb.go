//ready

package database

import (
	mongodb "citizenship/pkg/client/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBInfrastructure struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, Host, Port, Username, Password, Database, AuthSource string) (*MongoDBInfrastructure, error) {
	client, err := mongodb.NewClient(ctx, Host, Port, Username, Password, Database, AuthSource)
	if err != nil {
		return nil, err
	}
	return &MongoDBInfrastructure{db: client}, nil
}