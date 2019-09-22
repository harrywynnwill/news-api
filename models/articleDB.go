package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type ArticleRepo struct {
	gorm.Model
	Title       string
	Category    string
	Url         string
	Description string `gorm:"type:TEXT"`
	UrlToImage  sql.NullString
	Provider    Provider
	Date        int64
}
