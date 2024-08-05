package model

type LeagueResponse struct {
	Stages []Stage `json:"Stages"`
}

type Stage struct {
	Snm         string       `json:"Snm"`
	CompN       string       `json:"CompN"`
	CompId      string       `json:"CompId"`
	Cnm         string       `json:"Cnm"`
	Sid         string       `json:"Sid"`
	Events      []TeamEvents `json:"Events"`
	LeagueTable LeagueTable  `json:"LeagueTable"`
}

type LeagueTable struct {
	L []League `json:"L"`
}

type League struct {
	Tables []Table `json:"Tables"`
}

type Table struct {
	Teams []TeamDetails `json:"team"`
}

type TeamDetails struct {
	Tnm string `json:"Tnm"`
	Tid string `json:"Tid"`
}

type CustomLeagueResponse struct {
	Name        string                      `json:"name"`
	ExternalId  string                      `json:"externalId"`
	Country     string                      `json:"country"`
	GroupLeague []CustomLeagueGroupResponse `json:"groupLeague"`
}

type CustomLeagueGroupResponse struct {
	Name       string             `json:"name"`
	ExternalId string             `json:"externalId"`
	Teams      []CustomLeagueTeam `json:"teams"`
}

type CustomLeagueTeam struct {
	Name       string `json:"name"`
	ExternalID string `json:"externalId"`
}

type CustomLeagueMatchesResponse struct {
	TournamentId string                `json:"tournamentId"`
	GroupId      string                `json:"groupId"`
	Matches      []CustomMatchResponse `json:"matches"`
}
