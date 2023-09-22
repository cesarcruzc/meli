package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	dbInstance    *mongo.Database
	instanceError error
	logger        *log.Logger
)

// GetMongoClient returns a mongo client instance
func GetMongoClient() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		instanceError = err
	}

	dbName := os.Getenv("MONGO_DB")

	dbInstance = client.Database(dbName)

	return dbInstance, instanceError
}

func InitLogger() *log.Logger {

	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	return logger
}
