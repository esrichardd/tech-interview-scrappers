package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/esrichardd/tech-interview-scrappers/internal/domain/model"
)

const baseURL = "https://prod-public-api.livescore.com/v1/api/app/stage/soccer/italy/serie-a/-5"

type LeagueService struct {
	BaseService
}

func (s *LeagueService) GetLeague() (model.CustomLeagueResponse, error) {
	leagueResponse, err := s.fetchAndDecodeLeagueResponse()
	if err != nil {
		return model.CustomLeagueResponse{}, err
	}

	if err := validateLeagueResponse(leagueResponse); err != nil {
		return model.CustomLeagueResponse{}, err
	}

	return processLeagueResponse(leagueResponse), nil
}

func (s *LeagueService) GetLeagueMatches() (model.CustomLeagueMatchesResponse, error) {
	leagueResponse, err := s.fetchAndDecodeLeagueResponse()
	if err != nil {
		return model.CustomLeagueMatchesResponse{}, err
	}

	return processLeagueMatchesResponse(leagueResponse), nil
}

func (s *LeagueService) fetchAndDecodeLeagueResponse() (model.LeagueResponse, error) {
	resp, err := s.FetchData(baseURL)
	if err != nil {
		return model.LeagueResponse{}, err
	}

	var leagueResponse model.LeagueResponse
	if err := s.DecodeResponse(resp, &leagueResponse); err != nil {
		return model.LeagueResponse{}, err
	}

	return leagueResponse, nil
}

func validateLeagueResponse(leagueResponse model.LeagueResponse) error {
	if len(leagueResponse.Stages) == 0 {
		return errors.New("no stages found in response")
	}

	stage := leagueResponse.Stages[0]
	if len(stage.LeagueTable.L) == 0 || len(stage.LeagueTable.L[0].Tables) == 0 || len(stage.LeagueTable.L[0].Tables[0].Teams) < 2 {
		return errors.New("insufficient data in response")
	}

	return nil
}

func processLeagueResponse(leagueResponse model.LeagueResponse) model.CustomLeagueResponse {
	stage := leagueResponse.Stages[0]
	teams := stage.LeagueTable.L[0].Tables[0].Teams

	customTeams := make([]model.CustomLeagueTeam, len(teams))
	for i, team := range teams {
		customTeams[i] = model.CustomLeagueTeam{
			Name:       team.Tnm,
			ExternalID: team.Tid,
		}
	}

	customGroups := []model.CustomLeagueGroupResponse{
		{
			Name:       stage.Snm,
			ExternalId: stage.Sid,
			Teams:      customTeams,
		},
	}

	return model.CustomLeagueResponse{
		Name:        stage.CompN,
		ExternalId:  stage.CompId,
		Country:     stage.Cnm,
		GroupLeague: customGroups,
	}
}

func processLeagueMatchesResponse(leagueResponse model.LeagueResponse) model.CustomLeagueMatchesResponse {
	stage := leagueResponse.Stages[0]
	matches := stage.Events

	customMatches := make([]model.CustomMatchResponse, len(matches))
	for i, match := range matches {

		formatedDate, err := ConvertToDate(match.Esd)
		if err != nil {
			fmt.Println("Error al convertir la fecha:", err)
			return model.CustomLeagueMatchesResponse{}
		}

		customMatches[i] = model.CustomMatchResponse{
			LocalTeam:             match.T1[0].Nm,
			VisitorTeam:           match.T2[0].Nm,
			ExternalLocalTeamID:   match.T1[0].ID,
			ExternalVisitorTeamID: match.T2[0].ID,
			Status:                match.Eps,
			Date:                  formatedDate,
			MatchId:               match.Eid,
		}
	}

	return model.CustomLeagueMatchesResponse{
		TournamentId: stage.CompId,
		GroupId:      stage.Sid,
		Matches:      customMatches,
	}
}

// ConvertToDate convierte una cadena en formato "yyyyMMddHHmmss" a un objeto time.Time
func ConvertToDate(input int) (time.Time, error) {
	// Convertir el entero a una cadena
	inputStr := strconv.FormatInt(int64(input), 10)

	// Verificar que la longitud de la cadena sea de 14 caracteres
	if len(inputStr) != 14 {
		return time.Time{}, fmt.Errorf("la entrada debe tener exactamente 14 dígitos en el formato 'yyyyMMddHHmmss'")
	}

	// Parsear la fecha y hora de la cadena
	t, err := time.Parse("20060102150405", inputStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error al parsear la fecha: %v", err)
	}

	return t, nil
}
