package articles

import (
	"gin-rest/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleRegister(router *gin.RouterGroup){
	router.POST("/", ArticleCreate)
}

func ArticleCreate(c *gin.Context)  {
	articleModelValidator := NewArticleValidator()
	if err := articleModelValidator.Bind(c); err != nil{
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&articleModelValidator.articleModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "created"})
}