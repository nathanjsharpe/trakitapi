package trakitapi

import (
	"net/http"
)

const customerBasePath = "customer"

type CustomerService interface {
	List() ([]Customer, *http.Response, error)
}

type Customer struct {
	Id          int             `json:"cid"`
	CompanyName string          `json:"companyname"`
	Account     CustomerAccount `json:"account"`
	Data        interface{}     `json:"data"`
}

type CustomerAccount struct {
	MaxAssetCount    interface{} `json:"maxAssetCount"`
	PhysicalAddress  string      `json:"physical_address"`
	Phone            string      `json:"phone"`
	CompanyName      string      `json:"companyName"`
	Description      string      `json:"description"`
	EnabledProtocols []string    `json:"enabled_protocols"`
	EnabledMaps      []string    `json:"enabled_maps"`
	Email            string      `json:"email"`
}

type CustomerServiceOp struct {
	client *Client
}

func (s *CustomerServiceOp) List() ([]Customer, *http.Response, error) {
	req, err := s.client.NewRequest("GET", customerBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	customers := make([]Customer, 0)
	resp, err := s.client.Do(req, &customers)
	if err != nil {
		return nil, resp, err
	}

	return customers, resp, err
}
