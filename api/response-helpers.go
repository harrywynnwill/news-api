package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(deserializedPayload interface{}, w http.ResponseWriter, statusCode int) {
	payload, err := json.Marshal(deserializedPayload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}