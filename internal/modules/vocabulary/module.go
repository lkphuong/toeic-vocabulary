package vocabulary

import (
	"os"

	"github.com/lkphuong/toeic-vocabulary/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

func init() {
	client := database.New()
	if client == nil {
		panic("Failed to initialize MongoDB client")
	}
	dbName := os.Getenv("DB_NAME")
	collection = database.MongoClient.Database(dbName).Collection("vocabulary")
}
