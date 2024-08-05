package api

import (
	"encoding/json"
	"net/http"

	"github.com/esrichardd/tech-interview-scrappers/internal/application/service"
	"github.com/gorilla/mux"
)

var matchService = service.MatchService{}

func GetIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	scorebard, err := matchService.GetIncidents(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scorebard)
}
