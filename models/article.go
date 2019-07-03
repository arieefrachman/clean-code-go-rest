package models

import (
	"gin-rest/common"
	"github.com/jinzhu/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body string `gorm:"size:255"`
	FirstName string `gorm:"size:255"`
	NumberCol int8
}

func (ArticleModel) TableName() string{
	return "articles"
}

type ArticleRequest struct {
	Title string `form:"title" json:"title" binding:"required"`
	Body string `form:"body" json:"body" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	NumberCol int8 `form:"number_col" json:"number_col" binding:"required"`
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
