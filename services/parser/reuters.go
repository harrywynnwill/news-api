package parser

import (
	"encoding/xml"
	"esqimo-news-app/models"
	"time"
)

type reutersParseService struct {}

func newReutersParseService() *reutersParseService {
	return &reutersParseService{}
}

var ReutersParseService = newReutersParseService()

func (p reutersParseService) Parse(rawNews []byte) (articles []*models.Article, err error) {
	reutersDTO := &models.ReutersDTO{}
	e := xml.Unmarshal(rawNews, reutersDTO)
	if e != nil {
		return nil, e
	}
	if err != nil {
		return nil, err
	}
	articles, err = reutersDTOToArticles(reutersDTO)
	return
}

func reutersDTOToArticles(dto *models.ReutersDTO) ([]*models.Article, error) {
	var articles = make([]*models.Article, len(dto.Channel.Item))
	if dto.Channel.Item != nil {
		for i, item := range dto.Channel.Item {
			parsedDate, e := time.Parse(time.RFC1123Z, item.PubDate)
			if e != nil {
				return nil, e
			}
			article := &models.Article{
				ArticleSummary: models.ArticleSummary{
					Title:      item.Title,
					Category:   item.Category, // TODO CATEGORY MAPPING
					Date:       parsedDate,
					UrlToImage: dto.Channel.Image.URL, // There is no per article image using the generic Reuters image
					Provider:   models.REUTERS ,
				},
				Url:         item.Link,
				Description: item.Description,
			}
			articles[i] = article
		}
	}
	return articles, nil
}
