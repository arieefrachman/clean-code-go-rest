package usecase

import (
	"context"
	"gin-rest/articles"
	"gin-rest/models"
	"time"
)

type articleUsecase struct {
	articleRepo articles.Repository
	contextTimeout time.Duration
}

func NewArticleUsecase(a articles.Repository, timeout time.Duration) articles.Usecase  {
	return &articleUsecase{
		a,
		timeout,
	}
}

func (a *articleUsecase) Store(ctx context.Context, m *models.ArticleRequest) error  {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	article := &models.ArticleModel{
		Title: m.Title,
		Body: m.Body,
		FirstName: m.FirstName,
		NumberCol: m.NumberCol,
	}

	err := a.articleRepo.Store(ctx, article)
	if err != nil{
		return err
	}
	return nil
}
