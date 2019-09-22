package parser

import (
	"esqimo-news-app/models"
)

type ParseService interface {
	Parse(rawNews []byte) ([]*models.Article, error)
}

