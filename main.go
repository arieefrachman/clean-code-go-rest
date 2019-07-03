package main

import (
	"gin-rest/articles/delivery/http"
	"gin-rest/articles/repository"
	"gin-rest/articles/usecase"
	"gin-rest/common"
	"gin-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func Migrate(db *gorm.DB){
	db.AutoMigrate(models.ArticleModel{})
}
func main() {
	db := common.InitDB()
	Migrate(db)
	defer db.Close()
	r := gin.Default()
	v1 := r.Group("/api")
	ar := repository.NewPgsqlArticleRepository(db)
	timeoutContext := time.Second
	au := usecase.NewArticleUsecase(ar, timeoutContext)
	http.ArticleRegister(v1.Group("/article"), au)

	r.Run("0.0.0.0:3010")
}
