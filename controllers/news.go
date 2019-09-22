package controllers

import (
	"esqimo-news-app/api"
	"esqimo-news-app/models"
	"esqimo-news-app/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func LoadNews(w http.ResponseWriter, r *http.Request) {
	log.Println("Load news from sources")
	allArticles := []*models.Article{}
	reutersArticles, e := services.ReutersProviderService.GetNews()
	allArticles = append(allArticles, reutersArticles...)
	if e != nil {
		log.Println("Error getting news from reuters", e)
		w.WriteHeader(http.StatusInternalServerError)
	}
	bbcArticles, e := services.BBCProviderService.GetNews()
	if e != nil {
		log.Println("Error getting news from bbc", e)
		w.WriteHeader(http.StatusInternalServerError)
	}
	allArticles = append(allArticles, bbcArticles...)
	e = services.NewsService.LoadNews(allArticles)
	if e != nil {
		log.Println("Error loading news", e)
		w.WriteHeader(http.StatusInternalServerError)
	}
	api.SendResponse(nil, w, http.StatusOK)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	log.Println("Get news")
	category := r.URL.Query().Get("category")
	provider := r.URL.Query().Get("provider")

	var articleSummaries []*models.ArticleSummary
	var e error

	articleSummaries, e = services.NewsService.GetNewsByQuery(category, provider)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	api.SendResponse(models.ArticleList{Articles: articleSummaries}, w, http.StatusOK)
}

func GetNewsByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Get news by ID")
	params := mux.Vars(r)
	idInt64, e := strconv.ParseUint(params["ID"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	article, e := services.NewsService.GetArticle(uint(idInt64))
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	api.SendResponse(article, w, http.StatusOK)
}
