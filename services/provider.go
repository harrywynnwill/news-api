package services

import (
	"esqimo-news-app/models"
	"esqimo-news-app/repository"
	"esqimo-news-app/services/parser"
	"io/ioutil"
	"net/http"
)

type ProviderService interface {
	GetNews() ([]*models.ArticleSummary, error)
}

type providerServiceImpl struct {
	provider          models.Provider
	sourcesRepository repository.SourcesRepository
	parseService      parser.ParseService // Maybe use reflection here to get the type - remover the provider from the constructor?
}

func newProviderService(sourcesRepository repository.SourcesRepository, parseService parser.ParseService,
	provider models.Provider) providerServiceImpl {
	return providerServiceImpl{
		provider:          provider,
		sourcesRepository: sourcesRepository,
		parseService:      parseService,
	}
}

var ReutersProviderService = newProviderService(repository.NewMysqlSourcesRepository(), parser.ReutersParseService, models.REUTERS)
var BBCProviderService = newProviderService(repository.NewMysqlSourcesRepository(), parser.BBCParseService, models.BBC)

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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
