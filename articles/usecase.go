package articles

import (
	"context"
	"gin-rest/models"
)

type Usecase interface {
	Store(ctx context.Context, a *models.ArticleRequest) error
} 
