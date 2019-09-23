package models

import (
	"time"
)

type ArticleSummary struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Category   string    `json:"category"`
	UrlToImage string    `json:"urlToImage"`
	Date       time.Time `json:"date"`
	Provider   Provider  `json:"provider"`
}

// A more detailed article when the user drills in from the list view
type Article struct {
	ArticleSummary
	Url         string `json:"url"`
	Description string `json:"description"`
}

type ArticleList struct {
	Articles Articles `json:"articles"`
	Meta
}

func NewArticleList(articles []*ArticleSummary, meta *Meta) *ArticleList {
	paginatedArticles := &ArticleList{
		Articles: articles,
		Meta: Meta{
			PageSize:     meta.PageSize,
			Offset:       meta.Offset,
			TotalRecords: meta.TotalRecords,
		},
	}
	return paginatedArticles
}

type Articles []*ArticleSummary

func (a Articles) Len() int {
	return len(a)
}
func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Articles) Less(i, j int) bool {
	return a[i].Date.After(a[j].Date)
}
