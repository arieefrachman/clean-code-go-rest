package articles

import (
	"context"
	"gin-rest/models"
)

type Repository interface {
	Store(ctx context.Context, a *models.ArticleModel) error
}
