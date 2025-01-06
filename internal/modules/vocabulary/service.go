package vocabulary

import (
	"context"
	"log"

	"github.com/lkphuong/toeic-vocabulary/internal/models"
	"github.com/lkphuong/toeic-vocabulary/internal/utils"
)

var (
	repository Repository
)

type Service struct{}

func init() {
	repository = Repository{}
}

func (s *Service) GetVocabularyPaginate(ctx context.Context, page int, limit int) *utils.Response {

	vocab, err := repository.GetVocabularyPaginate(ctx, page, limit)

	if err != nil {
		return utils.NewResponse(nil, "Error while fetching vocabulary", 500)
	}

	totalCount, err := repository.CountVocabulary(ctx)

	if err != nil {
		return utils.NewResponse(nil, "Error while counting vocabulary", 500)
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	response := &PaginatedResult[models.Vocabulary]{
		Data:       vocab,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}

	return utils.NewResponse(response, "Vocabulary fetched successfully", 200)
}

func (s *Service) SaveVocabulary(ctx context.Context, vocab models.Vocabulary) *utils.Response {

	isExist, err := repository.GetVocabularyByWord(ctx, vocab.Word)

	log.Println("isExist", isExist.Word)

	if isExist.Word != "" {
		return utils.NewResponse(nil, "Vocabulary already exists", 400)
	}

	if err != nil {
		return utils.NewResponse(nil, "Error while checking vocabulary", 500)
	}

	err = repository.CreateVocabulary(ctx, vocab)

	if err != nil {
		return utils.NewResponse(nil, "Error while saving vocabulary", 500)
	}

	return utils.NewResponse(nil, "Vocabulary saved successfully", 200)
}
