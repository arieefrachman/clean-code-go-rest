package http

import (
	"gin-rest/articles"
	"gin-rest/common"
	"gin-rest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleHandler struct {
	AUsecase articles.Usecase
}

func ArticleRegister(router *gin.RouterGroup, usecase articles.Usecase){
	handler := &ArticleHandler{
		usecase,
	}
	router.POST("/", handler.ArticleCreate)
}

func (a *ArticleHandler) ArticleCreate(c *gin.Context)  {
	request := models.ArticleRequest{}
	if err := c.BindJSON(&request); err != nil{
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	a.AUsecase.Store(c, &request)

	c.JSON(http.StatusCreated, gin.H{"msg": request})
}


