package api

import (
	"encoding/json"
	"net/http"

	"github.com/esrichardd/tech-interview-scrappers/internal/application/service"
)

var leagueService = service.LeagueService{}

func GetLeagueHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := leagueService.GetLeague()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func GetLeagueMatchesHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := leagueService.GetLeagueMatches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
