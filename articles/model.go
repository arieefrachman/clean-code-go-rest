package articles

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

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
