package vocabulary

import "github.com/gin-gonic/gin"

func VocabularyRouter(r *gin.RouterGroup) {
	vocab := r.Group("/vocabulary")
	{
		vocab.GET("/", GetVocabularyPaginate)
		vocab.POST("/", CreateVocabulary)
	}
}
