package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/initialize", InitializeDataHandler).Methods("POST")
	router.HandleFunc("/league", GetLeagueHandler).Methods("GET")
	router.HandleFunc("/league/matches", GetLeagueMatchesHandler).Methods("GET")
	router.HandleFunc("/team/{id:[0-9]+}/matches", GetMatchesHandler).Methods("GET")
	router.HandleFunc("/match/{id:[0-9]+}/scorebard", GetIncidentsHandler).Methods("GET")
	return router
}

func StartServer() {
	router := NewRouter()
	PORT := os.Getenv("TECH_INTERVIEW_SCRAPPERS_PORT")

	if PORT == "" {
		PORT = "3003"
	}

	log.Println("Server starting at port", PORT)
	http.ListenAndServe(":"+PORT, router)
}
