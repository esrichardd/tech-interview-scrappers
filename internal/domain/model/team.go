package model

import "time"

type TeamResponse struct {
	Stages []TeamStage `json:"Stages"`
	Nm     string      `json:"Nm"`
	ID     string      `json:"ID"`
}

type TeamStage struct {
	Snm    string       `json:"Snm"`
	Events []TeamEvents `json:"Events"`
}

type TeamEvents struct {
	T1  []MatchTeamDetails `json:"T1"`
	T2  []MatchTeamDetails `json:"T2"`
	Eid string             `json:"Eid"`
	Eps string             `json:"Eps"`
	Esd int                `json:"Esd"`
	Sid string             `json:"Sid"`
}

type MatchTeamDetails struct {
	Nm string `json:"Nm"`
	ID string `json:"ID"`
}

type CustomTeamResponse struct {
	TeamName   string                `json:"teamName"`
	ExternalID string                `json:"ExternalId"`
	Matches    []CustomMatchResponse `json:"matches"`
}

type CustomMatchResponse struct {
	LocalTeam             string    `json:"localTeam"`
	VisitorTeam           string    `json:"visitorTeam"`
	ExternalLocalTeamID   string    `json:"externalLocalTeamId"`
	ExternalVisitorTeamID string    `json:"externalVisitorTeamId"`
	Status                string    `json:"status"`
	Date                  time.Time `json:"date"`
	LeagueID              string    `json:"leagueId"`
	MatchId               string    `json:"matchId"`
}
