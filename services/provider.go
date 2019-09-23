package services

import (
	"esqimo-news-app/models"
	"esqimo-news-app/repositories"
	"esqimo-news-app/services/parser"
	"io/ioutil"
	"net/http"
)

type ProviderService interface {
	GetNews() ([]*models.ArticleSummary, error)
}

type providerServiceImpl struct {
	provider          models.Provider
	sourcesRepository repositories.SourcesRepository
	parseService      parser.ParseService // Maybe use reflection here to get the type - remove the provider from the constructor?
}

func newProviderService(sourcesRepository repositories.SourcesRepository, parseService parser.ParseService,
	provider models.Provider) providerServiceImpl {
	return providerServiceImpl{
		provider:          provider,
		sourcesRepository: sourcesRepository,
		parseService:      parseService,
	}
}

var ReutersProviderService = newProviderService(repositories.MySqlSourcesRepository, parser.ReutersParseService, models.REUTERS)
var BBCProviderService = newProviderService(repositories.MySqlSourcesRepository, parser.BBCParseService, models.BBC)

func (p providerServiceImpl) GetNews() ([]*models.Article, error) {
	sources, e := p.sourcesRepository.GetBySource(p.provider)
	if e != nil {
		return nil, e
	}
	var allArticles []*models.Article
	for _, s := range sources {
		rawNews, e := requestNews(s.SourceUrl)
		if e != nil {
			return nil, e
		}
		articles, e := p.parseService.Parse(rawNews)
		if e != nil {
			return nil, e
		}
		allArticles = append(articles, allArticles...)
	}

	return allArticles, e
}

func requestNews(url string) ([]byte, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	body, e := ioutil.ReadAll(resp.Body)
	return body, e
}
