package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/esrichardd/tech-interview-scrappers/internal/application/service"
	"github.com/gorilla/mux"
)

var teamService = service.TeamService{}

func GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid Team ID", http.StatusBadRequest)
		return
	}

	matches, err := teamService.GetMatches(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}
