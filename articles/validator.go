package articles

import (
	"gin-rest/common"
	"github.com/gin-gonic/gin"
)

type ArticleModelValidator struct {
	Article struct{
		Title string `form:"title" json:"title"`
		Body string `form:"body" json:"body"`
		FirstName string `form:"first_name" json:"first_name"`
		NumberCol int8 `form:"number_col" json:"number_col"`
	}
	articleModel ArticleModel `json:"-"`
}

func NewArticleValidator() ArticleModelValidator  {
	return ArticleModelValidator{}
}

func (s *ArticleModelValidator) Bind(c *gin.Context) error  {
	err := common.Bind(c, s)

	if err != nil {
		return err
	}

	s.articleModel.Title = s.Article.Title
	s.articleModel.Body = s.Article.Body
	s.articleModel.FirstName = s.Article.FirstName
	s.articleModel.NumberCol = s.Article.NumberCol

	return nil
}
