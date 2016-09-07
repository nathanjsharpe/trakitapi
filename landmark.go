package trakitapi

import (
	"net/http"
)

const landmarkBasePath = "landmark"

type LandmarksService interface {
	List() ([]Landmark, *http.Response, error)
}

type Landmark struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Description string  `json:"desciption"`
	Icon        string  `json:"icon"`
	ShowOnMap   bool    `json:"showOnMap"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Status      string  `json:"status"`
}

type LandmarkServiceOp struct {
	client *Client
}

func (s *LandmarkServiceOp) list(path string) ([]Landmark, *http.Response, error) {
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	landmarks := make([]Landmark, 0)
	resp, err := s.client.Do(req, &landmarks)
	if err != nil {
		return nil, resp, err
	}

	return landmarks, resp, err
}

func (s *LandmarkServiceOp) List() ([]Landmark, *http.Response, error) {
	return s.list(landmarkBasePath)
}
