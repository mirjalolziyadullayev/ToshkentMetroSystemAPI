package Handlers

import (
	"Toshkent_Metro_System/models"
	"encoding/json"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Get(w, r)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	var CitiesData []models.Model
	ByteCollection, _ := os.ReadFile("db/all.json")
	json.Unmarshal(ByteCollection, &CitiesData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CitiesData)
}
