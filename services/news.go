package services

import (
	"esqimo-news-app/models"
	"esqimo-news-app/repository"
)

type NewService interface {
	LoadNews(articleList []*models.Article) error
	GetNews(offset, pageSize int) (*models.ArticleList, error)
	GetNewsByQuery(category, provider string, offset, pageSize int) (*models.ArticleList, error)
	GetArticle(id uint) (*models.Article, error)
}

type newsServiceImpl struct {
	articleRepository repository.ArticleRepository
}

func newNewsService(repo repository.ArticleRepository) newsServiceImpl {
	return newsServiceImpl{
		articleRepository: repo,
	}
}

var NewsService = newNewsService(repository.MysqlArticleRepository) // Singleton

func (n newsServiceImpl) LoadNews(articles []*models.Article) error {
	return n.articleRepository.Create(articles)
}

func (n newsServiceImpl) GetNews(offset, pageSize int) (*models.ArticleList, error) {
	articles, meta, e := n.articleRepository.Get(offset, pageSize)
	if e != nil {
		return nil, e
	}
	paginatedArticles := models.NewArticleList(articles, meta)
	return paginatedArticles, nil
}

func (n newsServiceImpl) GetNewsByQuery(category, provider string, offset, pageSize int) (*models.ArticleList, error) {
	articles, meta, e := n.articleRepository.GetByQuery(category, provider, offset, pageSize)
	if e != nil {
		return nil, e
	}
	paginatedArticles := models.NewArticleList(articles, meta)
	return paginatedArticles, nil
}

func (n newsServiceImpl) GetArticle(id uint) (*models.Article, error) {
	article, e := n.articleRepository.GetByID(id)
	if e != nil {
		return nil, e
	}
	return article, nil
}
