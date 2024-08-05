package ports

import "github.com/esrichardd/tech-interview-scrappers/internal/domain/model"

type TeamServicePort interface {
	GetMatches(id string) (model.CustomTeamResponse, error)
}

type LeagueServicePort interface {
	GetLeague() (model.CustomLeagueResponse, error)
}

type MatchServicePort interface {
	GetIncidents(id string) (model.CustomMatchIncidents, error)
}
