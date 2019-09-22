package main

import (
	"esqimo-news-app/controllers"
	"esqimo-news-app/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Query string

const (
	CATEGORY Query = "category"
	PROVIDER Query = "provider"
)

func main() {
	database.InitDB()
	port := "8000"
	r := mux.NewRouter()
	r.HandleFunc("/news", controllers.LoadNews).Methods("POST") // Would not be part of the public API
	r.HandleFunc("/news", controllers.GetNews).Methods("GET")
	r.HandleFunc("/news/{ID:[0-9]+}", controllers.GetNewsByID).Methods("GET")

	log.Println("Listening on port " + port + "...")

	log.Fatal(http.ListenAndServe(":"+port, r))

}
