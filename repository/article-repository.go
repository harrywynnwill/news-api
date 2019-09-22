package repository

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

type MysqlArticleRepository struct{}

func NewMysqlArticleRepository() *MysqlArticleRepository {
	return &MysqlArticleRepository{}
}

func (r MysqlArticleRepository) Create(articles []*models.Article) error {
	var e error
	database.WithDb(func(db *gorm.DB) {
		articleRepos := toArticleReposFromArticle(articles)
		e = gormbulk.BulkInsert(db, articleRepos, 3000)
	})
	return e
}

func (r MysqlArticleRepository) Get(offset, pageSize int) ([]*models.ArticleSummary, *models.Meta, error) {
	var e error
	var articleRepo []*models.ArticleRepo
	var totalRecords uint
	database.WithDb(func(db *gorm.DB) {
		e = db.Debug().Model(models.ArticleRepo{}).
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

func (r MysqlArticleRepository) GetByID(id uint) (*models.Article, error) {
	var e error
	var articleRepo models.ArticleRepo
	database.WithDb(func(db *gorm.DB) {
		e = db.Model(models.ArticleRepo{}).First(&articleRepo, id).Error
	})
	if e != nil {
		return nil, e
	}

	enrichedArticle := &models.Article{}
	toArticleFromArticleRepo(enrichedArticle, &articleRepo)
	return enrichedArticle, nil
}

func (r MysqlArticleRepository) GetByQuery(category, provider string, offset, pageSize int) ([]*models.ArticleSummary, *models.Meta, error) {
	var e error
	var articleRepo []*models.ArticleRepo
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
		e = db.Debug().
			Model(models.ArticleRepo{}).
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
	var articleRepos = make([]interface{}, length)
	for i, a := range articles {
		var articleRepo = models.ArticleRepo{}
		articleRepo.Date = a.Date.Unix()
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

func toEnrichedArticlesFromArticleRepo(articleRepo *models.ArticleRepo) (*models.Article, error) {
	var article = &models.Article{}
	toArticleFromArticleRepo(article, articleRepo)
	return article, nil
}

func toArticlesFromArticleRepo(ar []*models.ArticleRepo) []*models.ArticleSummary {
	var articles = make([]*models.ArticleSummary, len(ar))
	for i, a := range ar {
		articles[i] = toArticleSummaryFromArticleRepo(a)
	}
	return articles
}

func toArticleSummaryFromArticleRepo(ar *models.ArticleRepo) *models.ArticleSummary {
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

func toArticleFromArticleRepo(article *models.Article, ar *models.ArticleRepo) {
	article.ArticleSummary = *toArticleSummaryFromArticleRepo(ar)
	article.Url = ar.Url
	article.Description = ar.Description
}