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

	var e error
	var articleList *models.ArticleList

	isValidQueryParams := api.IsValidQueryParams(category, provider)
	if !isValidQueryParams {
		w.WriteHeader(http.StatusBadRequest)
	}

	pageSizeInt, e := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	offSetInt, e := strconv.Atoi(r.URL.Query().Get("offset"))
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if pageSizeInt > 300 || offSetInt > 300 {
		w.WriteHeader(http.StatusBadRequest)
	}

	if category != "" || provider != "" {
		articleList, e = services.NewsService.GetNewsByQuery(category, provider, offSetInt, pageSizeInt)
	} else {
		articleList, e = services.NewsService.GetNews(offSetInt, pageSizeInt)
	}
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	api.SendResponse(articleList, w, http.StatusOK)
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
