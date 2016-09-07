package trakitapi

import (
	"errors"
	"net/http"
)

const sessionTokenBasePath = "session_token"

type SessionTokenService interface {
	Create(*SessionTokenCreateRequest) (*SessionToken, *User, *http.Response, error)
}

type SessionToken struct {
	Token      string `json:"token"`
	Tenant     string `json:"tenant"`
	Expiration string `json:"expiration"`
	Level      int    `json:"level"`
}

func (st *SessionToken) IsExpired() bool {
	return false
}

type SessionTokenServiceOp struct {
	client *Client
}

type SessionTokenCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionTokenCreateResponse struct {
	Key     SessionToken `json:"key"`
	User    User         `json:"user"`
	Error   string       `json:"error"`
	Success bool
}

func (s *SessionTokenServiceOp) Create(createRequest *SessionTokenCreateRequest) (*SessionToken, *User, *http.Response, error) {
	if createRequest == nil {
		return nil, nil, nil, errors.New("createRequest cannot be nil")
	}

	path := sessionTokenBasePath

	req, err := s.client.NewRequest("POST", path, createRequest)
	if err != nil {
		return nil, nil, nil, err
	}

	var respContent SessionTokenCreateResponse
	resp, err := s.client.Do(req, &respContent)
	if err != nil {
		return nil, nil, resp, err
	}
	if respContent.Error != "" {
		return nil, nil, resp, errors.New(respContent.Error)
	}

	return &respContent.Key, &respContent.User, resp, err
}
