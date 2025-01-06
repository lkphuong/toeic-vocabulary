package vocabulary

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lkphuong/toeic-vocabulary/internal/models"
	"github.com/lkphuong/toeic-vocabulary/internal/utils"
)

var (
	service Service
)

func GetVocabularyPaginate(c *gin.Context) {

	r := c.Request

	ctx := r.Context()

	page, limit := utils.GetPaginationParams(r)

	result := service.GetVocabularyPaginate(ctx, page, limit)

	utils.JSONResponse(*result, c)
}

func CreateVocabulary(c *gin.Context) {

	r := c.Request

	ctx := r.Context()

	var vocab models.Vocabulary

	if err := c.ShouldBindBodyWithJSON(&vocab); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	vocab.CreatedAt = time.Now()

	result := service.SaveVocabulary(ctx, vocab)

	utils.JSONResponse(*result, c)
}
