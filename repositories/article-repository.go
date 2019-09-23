package repositories

import (
	"database/sql"
	"esqimo-news-app/database"
	"esqimo-news-app/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
	"time"
)

type ArticleRepository interface {
	Create(articleList []*models.Article) error
	Get(offset, pageSize int) ([]*models.ArticleSummary, *models.Meta, error)
	GetByQuery(category, provider string, offset, pageSize int) ([]*models.ArticleSummary, *models.Meta, error)
	GetByID(id uint) (*models.Article, error)
}

type mysqlArticleRepository struct{}

func newMysqlArticleRepository() *mysqlArticleRepository {
	return &mysqlArticleRepository{}
}

var MysqlArticleRepository = newMysqlArticleRepository()

func (r mysqlArticleRepository) Create(articles []*models.Article) error {
	var e error
	database.WithDb(func(db *gorm.DB) {
		articleRepos := toArticleReposFromArticle(articles)
		if articleRepos != nil {
			e = gormbulk.BulkInsert(db, articleRepos, 3000)
		}
	})
	return e
}

func (r mysqlArticleRepository) Get(offset, pageSize int) ([]*models.ArticleSummary, *models.Meta, error) {
	var e error
	var articleRepo []*models.ArticleRepository
	var totalRecords uint
	database.WithDb(func(db *gorm.DB) {
		e = db.Model(models.ArticleRepository{}).
			Offset(offset).
			Count(&totalRecords).
			Limit(pageSize).
			Order("date desc").
			Find(&articleRepo).
			Error
	})
	if e != nil {
		return nil, nil, e
	}
	articles := toArticlesFromArticleRepo(articleRepo)
	meta := models.NewMeta(pageSize, offset, totalRecords)
	return articles, meta, nil
}

func (r mysqlArticleRepository) GetByID(id uint) (*models.Article, error) {
	var e error
	var articleRepo models.ArticleRepository
	database.WithDb(func(db *gorm.DB) {
		e = db.Model(models.ArticleRepository{}).First(&articleRepo, id).Error
	})
	if e != nil {
		return nil, e
	}

	enrichedArticle := &models.Article{}
	toArticleFromArticleRepo(enrichedArticle, &articleRepo)
	return enrichedArticle, nil
}

func (r mysqlArticleRepository) GetByQuery(category, provider string, offset, pageSize int) ([]*models.ArticleSummary, *models.Meta, error) {
	var e error
	var articleRepo []*models.ArticleRepository
	var totalRecords uint
	// Need to use a map and not the repo struct as it will use zero values on the empty properties e.g. ID: O
	dbQuery := map[string]interface{}{}
	if category != "" {
		dbQuery["category"] = category
	}
	if provider != "" {
		dbQuery["provider"] = provider
	}

	database.WithDb(func(db *gorm.DB) {
		e = db.Model(models.ArticleRepository{}).
			Offset(offset).
			Count(&totalRecords).
			Limit(pageSize).
			Where(dbQuery).
			Order("date desc").
			Find(&articleRepo).Error
	})

	if e != nil {
		return nil, nil, e
	}
	articleSummaries := toArticlesFromArticleRepo(articleRepo)
	meta := models.NewMeta(pageSize, offset, totalRecords)
	return articleSummaries, meta, nil
}

func toArticleReposFromArticle(articles []*models.Article) []interface{} {
	length := len(articles)
	if length == 0 {
		return nil
	}
	var articleRepos = make([]interface{}, length)
	for i, a := range articles {
		var articleRepo = models.ArticleRepository{}
		if !a.Date.IsZero() {
			articleRepo.Date = a.Date.Unix()
		}
		articleRepo.Category = a.Category
		articleRepo.Url = a.Url
		articleRepo.UrlToImage = sql.NullString{String: a.UrlToImage, Valid: a.UrlToImage != ""}
		articleRepo.Description = a.Description
		articleRepo.Title = a.Title
		articleRepo.Provider = a.Provider
		articleRepos[i] = articleRepo
	}
	return articleRepos
}

func toArticlesFromArticleRepo(ar []*models.ArticleRepository) []*models.ArticleSummary {
	var articles = make([]*models.ArticleSummary, len(ar))
	for i, a := range ar {
		if articleRepo := toArticleSummaryFromArticleRepo(a); articleRepo != nil {
			articles[i] = articleRepo
		}
	}
	return articles
}

func toArticleSummaryFromArticleRepo(ar *models.ArticleRepository) *models.ArticleSummary {
	article := &models.ArticleSummary{}
	if ar.ID != 0 {
		article.ID = ar.ID
	}
	if ar.Title != "" {
		article.Title = ar.Title
	}
	if ar.Category != "" {
		article.Category = ar.Category
	}
	if ar.UrlToImage.Valid {
		article.UrlToImage = ar.UrlToImage.String
	}
	if ar.Date != 0 {
		unixDate := time.Unix(ar.Date, 0)
		article.Date = unixDate
	}
	if ar.Provider != "" {
		article.Provider = ar.Provider
	}

	return article
}

func toArticleFromArticleRepo(article *models.Article, ar *models.ArticleRepository) {
	if article != nil {
		if ar != nil {
			if articlesSummary := toArticleSummaryFromArticleRepo(ar); articlesSummary != nil {
				article.ArticleSummary = *articlesSummary
				if ar.Url != "" {
					article.Url = ar.Url
				}
				if ar.Description != "" {
					article.Description = ar.Description
				}
			}
		}

	}
}
