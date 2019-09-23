package models

import (
	"github.com/jinzhu/gorm"
)

type SourcesRepository struct {
	gorm.Model
	SourceUrl string
	Provider  Provider
}
