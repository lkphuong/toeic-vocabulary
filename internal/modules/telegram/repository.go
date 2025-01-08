package telegram

import (
	"context"
	"time"

	"github.com/lkphuong/toeic-vocabulary/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct{}

func (r *Repository) GetChatIDs(ctx context.Context) ([]models.ChatID, error) {

	var chats []models.ChatID

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &chats); err != nil {
		return nil, err
	}

	return chats, nil
}

func (r *Repository) GetByChatID(chatID int64) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chat *models.ChatID

	err := collection.FindOne(ctx, bson.M{"chat_id": chatID}).Decode(&chat)

	return err != mongo.ErrNoDocuments
}

func (r *Repository) Save(chatID int64, username string) error {

	var chat models.ChatID

	chat.ChatID = chatID
	chat.Username = username

	_, err := collection.InsertOne(context.Background(), chat)

	if err != nil {
		return err
	}

	return nil
}
