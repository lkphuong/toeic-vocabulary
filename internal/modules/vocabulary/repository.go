package vocabulary

import (
	"context"

	"github.com/lkphuong/toeic-vocabulary/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct{}

func (r *Repository) GetVocabularyPaginate(ctx context.Context, page int, limit int) ([]models.Vocabulary, error) {
	var vocab []models.Vocabulary

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &vocab); err != nil {
		return nil, err
	}

	return vocab, nil
}

func (r *Repository) CountVocabulary(ctx context.Context) (int64, error) {
	totalCount, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

func (r *Repository) GetVocabularyByWord(ctx context.Context, word string) (*models.Vocabulary, error) {
	var vocab models.Vocabulary

	err := collection.FindOne(context.Background(), bson.M{"word": word}).Decode(&vocab)

	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	return &vocab, nil
}

func (r *Repository) CreateVocabulary(ctx context.Context, vocab models.Vocabulary) error {
	_, err := collection.InsertOne(context.Background(), vocab)
	if err != nil {
		return err
	}

	return nil
}
