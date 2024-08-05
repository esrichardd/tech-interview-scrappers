package model

type MatchIncidentsResponse struct {
	Eid         string                  `json:"Eid"`
	Incs        Incs                    `json:"Incs"`
	LocalTeam   []TeamIncidentsResponse `json:"T1"`
	VisitorTeam []TeamIncidentsResponse `json:"T2"`
}

type MatchScorebardResponse struct {
	Eps   string                  `json:"Eps"`
	Team1 []TeamScorebardResponse `json:"T1"`
	Team2 []TeamScorebardResponse `json:"T2"`
}

type TeamScorebardResponse struct {
	Nm string `json:"Nm"`
	ID string `json:"ID"`
}

type TeamIncidentsResponse struct {
	Nm string `json:"Nm"`
	ID string `json:"ID"`
}

type Incs map[string][]IncDetail

type IncDetail struct {
	Min  int         `json:"Min"`
	Pn   string      `json:"Pn"`
	IT   int         `json:"IT"`
	Incs []IncDetail `json:"Incs,omitempty"`
}

type ResultDetail struct {
	Min     int            `json:"min"`
	Actions []ResultAction `json:"actions"`
}

type ResultAction struct {
	PlayerName string `json:"playerName"`
	Action     string `json:"action"`
}

type Result map[string][]ResultDetail

type CustomMatchIncidents struct {
	MatchId     string              `json:"matchId"`
	Status      string              `json:"status"`
	LocalTeam   CustomTeamIncidents `json:"localTeam"`
	VisitorTeam CustomTeamIncidents `json:"visitorTeam"`
	Scorebard   Result              `json:"scorebard"`
}

type CustomTeamIncidents struct {
	TeamId   string `json:"teamId"`
	TeamName string `json:"teamName"`
}

type MapActions map[int]string
