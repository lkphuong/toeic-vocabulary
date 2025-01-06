package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/lkphuong/toeic-vocabulary/internal/modules/vocabulary"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		vocabulary.VocabularyRouter(v1)
	}
}
