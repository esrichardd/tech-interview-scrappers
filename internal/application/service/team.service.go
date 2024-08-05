package service

import (
	"errors"
	"fmt"

	"github.com/esrichardd/tech-interview-scrappers/internal/domain/model"
)

type TeamService struct {
	BaseService
}

func (s *TeamService) GetMatches(id int) (model.CustomTeamResponse, error) {
	url := fmt.Sprintf("https://team-api.livescore.com/v1/api/app/team/%d/details?locale=en&MD=1", id)
	resp, err := s.FetchData(url)
	if err != nil {
		return model.CustomTeamResponse{}, err
	}

	var teamResponse model.TeamResponse
	if err := s.DecodeResponse(resp, &teamResponse); err != nil {
		return model.CustomTeamResponse{}, err
	}

	if err := validateTeamResponse(teamResponse); err != nil {
		return model.CustomTeamResponse{}, err
	}

	customResponse := processTeamResponse(teamResponse)
	return customResponse, nil
}

func validateTeamResponse(teamResponse model.TeamResponse) error {
	if len(teamResponse.Stages) == 0 {
		return errors.New("no stages found in response")
	}

	stage := teamResponse.Stages[0]
	if len(stage.Events) == 0 {
		return errors.New("no events found in response")
	}

	return nil
}

func processTeamResponse(teamResponse model.TeamResponse) model.CustomTeamResponse {
	stage := teamResponse.Stages[0]
	matches := stage.Events

	var customMatches []model.CustomMatchResponse
	for _, match := range matches {

		formatedDate, err := ConvertToDate(match.Esd)
		if err != nil {
			fmt.Println("Error al convertir la fecha:", err)
			return model.CustomTeamResponse{}
		}

		customMatches = append(customMatches, model.CustomMatchResponse{
			LocalTeam:             match.T1[0].Nm,
			VisitorTeam:           match.T2[0].Nm,
			ExternalLocalTeamID:   match.T1[0].ID,
			ExternalVisitorTeamID: match.T2[0].ID,
			Status:                match.Eps,
			Date:                  formatedDate,
			LeagueID:              match.Sid,
			MatchId:               match.Eid,
		})
	}

	return model.CustomTeamResponse{
		TeamName:   teamResponse.Nm,
		ExternalID: teamResponse.ID,
		Matches:    customMatches,
	}
}
