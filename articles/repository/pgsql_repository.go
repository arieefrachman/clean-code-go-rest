package repository

import (
	"context"
	"gin-rest/articles"
	"gin-rest/models"
	"github.com/jinzhu/gorm"
)

type pgsqlArticleRepository struct {
	Conn *gorm.DB
}

func NewPgsqlArticleRepository(Conn *gorm.DB) articles.Repository  {
	return &pgsqlArticleRepository{Conn}
}

func (p *pgsqlArticleRepository) Store(ctx context.Context, a *models.ArticleModel)  error {
	err := p.Conn.Save(&a).Error
	return err
}