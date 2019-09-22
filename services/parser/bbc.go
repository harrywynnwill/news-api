package parser

import (
	"encoding/xml"
	"esqimo-news-app/models"
	"strings"
	"time"
)

type bBCParseService struct {
	provider models.Provider
}

func newBBCParseService() *bBCParseService {
	return &bBCParseService{provider: models.BBC,}
}

var BBCParseService = newBBCParseService()

func (p bBCParseService) Parse(rawNews []byte) (articles []*models.Article, err error) {
	bBCDTO := &models.BBCDTO{}
	e := xml.Unmarshal(rawNews, bBCDTO)
	if e != nil {
		return nil, e
	}

	if err != nil {
		return nil, err
	}
	articles, err = bBCDTOToArticles(bBCDTO)
	return
}

func bBCDTOToArticles(dto *models.BBCDTO) ([]*models.Article, error) {
	var articles = make([]*models.Article, len(dto.Channel.Item))
	if dto.Channel.Item != nil {
		for i, item := range dto.Channel.Item {
			parsedDate, e := time.Parse(time.RFC1123, item.PubDate)
			if e != nil {
				return nil, e
			}
			article := &models.Article{
				ArticleSummary: models.ArticleSummary{
					Title:      strings.TrimSpace(item.Title),
					Category:   "", // TODO CATEGORY MAPPING
					Date:       parsedDate,
					UrlToImage: dto.Channel.Image.URL, //no images
					Provider:   models.BBC,
				},
				Url:         item.Link,
				Description: strings.TrimSpace(item.Description),
			}
			articles[i] = article
		}
	}
	// TODO implement this
	return articles, nil
}
