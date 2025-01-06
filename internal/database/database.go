package database

import (
	"context"
	"log"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri         = os.Getenv("DB_URI")
	MongoClient *mongo.Client
	once        sync.Once
)

func New() *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		MongoClient = client
		log.Println("Connected to MongoDB")
	})

	return MongoClient
}
