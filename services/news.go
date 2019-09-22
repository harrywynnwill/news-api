package services

import (
	"esqimo-news-app/models"
	"esqimo-news-app/repository"
	"sort"
)

type NewService interface {
	LoadNews(articleList []*models.Article) error
	GetNews() ([]*models.ArticleSummary, error)
	GetNewsByQuery(category, provider string) ([]*models.ArticleSummary, error)
	//GetNewsByProvider(provider string) ([]*models.ArticleSummary, error)
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

var NewsService = newNewsService(repository.NewMysqlArticleRepository()) // Singleton

func (n newsServiceImpl) LoadNews(articles []*models.Article) error {
	return n.articleRepository.Create(articles)
}

func (n newsServiceImpl) GetNews() ([]*models.ArticleSummary, error) {
	articles, e := n.articleRepository.Get()
	if e != nil {
		return nil, e
	}
	sort.Sort(models.Articles(articles))
	return articles, nil
}

func (n newsServiceImpl) GetNewsByQuery(category, provider string) (articles []*models.ArticleSummary, e error) {
	articles, e = n.articleRepository.GetByQuery(category, provider)
	return
}

//func (n newsServiceImpl) GetNewsByProvider(provider string) (articles []*models.ArticleSummary, e error) {
//	articles, e = n.articleRepository.GetByProvider(provider)
//	return
//}

func (n newsServiceImpl) GetArticle(id uint) (*models.Article, error) {
	article, e := n.articleRepository.GetByID(id)
	if e != nil {
		return nil, e
	}
	return article, nil
}
