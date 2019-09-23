package repositories

import (
	"esqimo-news-app/database"
	"esqimo-news-app/models"
	"github.com/jinzhu/gorm"
)

type SourcesRepository interface {
	GetBySource(provider models.Provider) ([]*models.SourcesRepository, error)
}

type mysqlSourcesRepository struct{}

func newMysqlSourcesRepository() *mysqlSourcesRepository {
	return &mysqlSourcesRepository{}
}

var MySqlSourcesRepository = newMysqlSourcesRepository()

func (r mysqlSourcesRepository) GetBySource(provider models.Provider) ([]*models.SourcesRepository, error) {
	var e error
	var sources []*models.SourcesRepository
	database.WithDb(func(db *gorm.DB) { e = db.Where(&models.SourcesRepository{Provider: provider}).Find(&sources).Error })
	return sources, e
}
