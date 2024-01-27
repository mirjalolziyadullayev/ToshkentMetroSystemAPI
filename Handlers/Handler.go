package Handlers

import (
	"Toshkent_Metro_System/models"
	"encoding/json"
	"net/http"
	"os"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetModels(w, r)
	}
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	var Data []models.Model
	ByteCollection, _ := os.ReadFile("db/all.json")
	json.Unmarshal(ByteCollection, &Data)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Data)
}
