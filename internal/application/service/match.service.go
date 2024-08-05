package service

import (
	"fmt"
	"sort"

	"github.com/esrichardd/tech-interview-scrappers/internal/domain/model"
)

type MatchService struct {
	BaseService
}

func (m *MatchService) GetIncidents(id string) (model.CustomMatchIncidents, error) {
	incidentsUrl := fmt.Sprintf("https://prod-public-api.livescore.com/v1/api/app/incidents/soccer/%s?locale=en", id)
	scorebardUrl := fmt.Sprintf("https://prod-public-api.livescore.com/v1/api/app/scoreboard/soccer/%s?locale=en", id)

	incidentsResp, err := m.FetchData(incidentsUrl)
	if err != nil {
		return model.CustomMatchIncidents{}, err
	}

	var matchIncidents model.MatchIncidentsResponse
	if err := m.DecodeResponse(incidentsResp, &matchIncidents); err != nil {
		return model.CustomMatchIncidents{}, err
	}

	scorebardResp, err := m.FetchData(scorebardUrl)
	if err != nil {
		return model.CustomMatchIncidents{}, err
	}

	var matchScorebard model.MatchScorebardResponse
	if err := m.DecodeResponse(scorebardResp, &matchScorebard); err != nil {
		return model.CustomMatchIncidents{}, err
	}

	result := m.processIncidents(matchIncidents)
	customMatchIncidents := m.createCustomMatchIncidents(id, matchScorebard, result)

	return customMatchIncidents, nil
}

func (m *MatchService) processIncidents(matchIncidents model.MatchIncidentsResponse) model.Result {
	mapActions := model.MapActions{
		36: "goal",
		37: "penal goal",
		43: "yellow card",
		44: "red card",
		63: "assistance",
	}

	result := make(model.Result)
	for team, actions := range matchIncidents.Incs {
		actionMap := make(map[int][]model.ResultAction)
		for _, action := range actions {
			m.processActions(action, mapActions, actionMap)
		}

		details := m.toResultDetails(actionMap)
		result[fmt.Sprintf("Time%s", team)] = details
	}

	return result
}

func (m *MatchService) processActions(inc model.IncDetail, mapActions model.MapActions, actionMap map[int][]model.ResultAction) {
	if inc.Incs != nil {
		for _, subInc := range inc.Incs {
			m.processActions(subInc, mapActions, actionMap)
		}
	}

	actionName := mapActions[inc.IT]

	if inc.Pn != "" && actionName != "" {
		action := model.ResultAction{
			PlayerName: inc.Pn,
			Action:     actionName,
		}
		actionMap[inc.Min] = append(actionMap[inc.Min], action)
	}
}

func (m *MatchService) toResultDetails(actionMap map[int][]model.ResultAction) []model.ResultDetail {
	details := []model.ResultDetail{}
	for min, actions := range actionMap {
		details = append(details, model.ResultDetail{
			Min:     min,
			Actions: actions,
		})
	}

	sort.Slice(details, func(i, j int) bool {
		return details[i].Min < details[j].Min
	})

	return details
}

func (m *MatchService) createCustomMatchIncidents(id string, matchScorebard model.MatchScorebardResponse, result model.Result) model.CustomMatchIncidents {
	return model.CustomMatchIncidents{
		MatchId:   id,
		Status:    matchScorebard.Eps,
		Scorebard: result,
		LocalTeam: model.CustomTeamIncidents{
			TeamId:   matchScorebard.Team1[0].ID,
			TeamName: matchScorebard.Team1[0].Nm,
		},
		VisitorTeam: model.CustomTeamIncidents{
			TeamId:   matchScorebard.Team2[0].ID,
			TeamName: matchScorebard.Team2[0].Nm,
		},
	}
}
