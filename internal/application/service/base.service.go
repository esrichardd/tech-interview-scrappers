package service

import (
	"encoding/json"
	"errors"
	"net/http"
)

type BaseService struct{}

func (s *BaseService) FetchData(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *BaseService) DecodeResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to get data")
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return err
	}

	return nil
}
