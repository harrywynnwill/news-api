package database

import (
	"esqimo-news-app/models"
	"esqimo-news-app/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type fn func(db *gorm.DB)

func WithDb(fun fn) error {
	settings.GetSettings()
	configStr := getConnectionString()
	db, err := gorm.Open("mysql", configStr)
	if err != nil {
		log.Println("Error starting DB", err)
	}
	fun(db)
	defer db.Close()
	return nil
}

func runMigrations() {
	WithDb(func(db *gorm.DB) {
		db.AutoMigrate(&models.ArticleRepo{})
		db.AutoMigrate(&models.SourcesRepository{})
	})
}

func addIndices() {
	WithDb(func(db *gorm.DB) {
		db.Model(&models.ArticleRepo{}).AddIndex("idx_date", "date")
	})
}

func InitDB() {
	runMigrations()
	addIndices()
}

func getConnectionString() string {
	appSettings := settings.GetSettings()
	sprintf := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		appSettings["databaseUser"],
		appSettings["databasePassword"],
		appSettings["databaseServer"],
		appSettings["databasePort"],
		appSettings["databaseName"])
	return sprintf
}
