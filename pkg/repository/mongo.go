package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const (
	usersCollection   = "users"
	companyCollection = "company"
)


func NewMongoDB(ctx context.Context, uri,dbName string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	logrus.Printf("URL to database %s", uri)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return db, nil
}