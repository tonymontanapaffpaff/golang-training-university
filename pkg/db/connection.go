package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(ctx context.Context, dbUser, dbPassword, dbHost, dbPort string) (*mongo.Client, error) {
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	//	fmt.Sprintf("mongodb://%v:%v@%v:%v/?sslmode=disable", dbUser, dbPassword, dbHost, dbPort)))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%v:%v", dbHost, dbPort)))
	if err != nil {
		return nil, fmt.Errorf("can't connect to database, error: %v", err)
	}
	return client, nil
}
