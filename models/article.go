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

type Articles []*ArticleSummary

type ArticleList struct {
	Articles Articles `json:"articles"`
	Meta
}

func (a Articles) Len() int {
	return len(a)
}
func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Articles) Less(i, j int) bool {
	return a[i].Date.After(a[j].Date)
}
