package vocabulary

import (
	"context"
	"fmt"
	"log"

	"github.com/lkphuong/toeic-vocabulary/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/rand"
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

func (r *Repository) GetRandomVocabulary(ctx context.Context) (*models.Vocabulary, error) {
	var vocab models.Vocabulary

	fmt.Println("Trying to get random vocabulary")
	// Get the total number of documents in the collection
	count, err := collection.CountDocuments(ctx, bson.D{})

	fmt.Println("Count:", count)
	fmt.Println("Error:", err)
	if err != nil {
		return nil, fmt.Errorf("failed to count documents: %w", err)
	}

	fmt.Println("Total number of documents:", count)

	// Nếu không có tài liệu nào trong cơ sở dữ liệu, trả về lỗi
	if count == 0 {
		return nil, mongo.ErrNoDocuments
	}

	// Randomly select an offset
	skip := rand.Int63n(count)

	fmt.Println("Skip:", skip)

	// Query with skip and limit 1 to get a random document
	cursor, err := collection.Find(ctx, bson.D{}, options.Find().SetSkip(skip).SetLimit(1))
	if err != nil {
		return nil, fmt.Errorf("failed to find random vocabulary: %w", err)
	}
	defer cursor.Close(ctx)

	// Always check if there is data in the cursor
	if cursor.Next(ctx) {
		if err := cursor.Decode(&vocab); err != nil {
			return nil, fmt.Errorf("failed to decode vocabulary: %w", err)
		}
	} else {
		// Return an empty vocabulary object if no document is found (which should not happen)
		log.Println("No document found, returning default vocabulary.")
		return &models.Vocabulary{}, nil
	}

	// Return the random vocabulary
	return &vocab, nil
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
