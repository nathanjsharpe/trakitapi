package trakitapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const versionsBasePath = "version/all"

const AppApiRegex = `trakitng_app_api-([^-]+)`
const DataApiRegex = `trakitng_data_api-([^-]+)`
const AuthApiRegex = `trakitng_auth_api-([^-]+)`
const EventApiRegex = `trakitng_event_api-([^-]+)`
const WebAppRegex = `Version:\D+(\d+\.\d+\.\d+)`

type VersionsService interface {
	List() (*Versions, *http.Response, error)
}

type VersionsServiceOp struct {
	client *Client
}

type VersionListResponse struct {
	App   string `json:"app-api"`
	Data  string `json:"data-api"`
	Auth  string `json:"auth-api"`
	Event string `json:"event-api"`
}

type Versions struct {
	App   string
	Data  string
	Auth  string
	Event string
	Web   string
}

func (s *VersionsServiceOp) List() (*Versions, *http.Response, error) {
	req, err := s.client.NewRequest("GET", versionsBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	var versionResponse VersionListResponse
	resp, err := s.client.Do(req, &versionResponse)
	if err != nil {
		return nil, resp, err
	}

	sanitizedVersions := sanitizeApiVersions(versionResponse)
	webAppVersion := getWebAppVersion(s.client.WebAppUrl)

	versions := Versions{
		App:   sanitizedVersions.App,
		Data:  sanitizedVersions.Data,
		Auth:  sanitizedVersions.Auth,
		Event: sanitizedVersions.Event,
		Web:   webAppVersion,
	}

	return &versions, resp, err
}

func sanitizeApiVersions(rawVersions VersionListResponse) VersionListResponse {
	appApiRegex, err := regexp.Compile(AppApiRegex)
	if err != nil {
		panic(err)
	}

	dataApiRegex, err := regexp.Compile(DataApiRegex)
	if err != nil {
		panic(err)
	}

	authApiRegex, err := regexp.Compile(AuthApiRegex)
	if err != nil {
		panic(err)
	}

	eventApiRegex, err := regexp.Compile(EventApiRegex)
	if err != nil {
		panic(err)
	}

	return VersionListResponse{
		App:   appApiRegex.FindStringSubmatch(rawVersions.App)[1],
		Data:  dataApiRegex.FindStringSubmatch(rawVersions.Data)[1],
		Auth:  authApiRegex.FindStringSubmatch(rawVersions.Auth)[1],
		Event: eventApiRegex.FindStringSubmatch(rawVersions.Event)[1],
	}
}

func getWebAppVersion(baseUrl *url.URL) string {
	p, _ := url.Parse("version")
	path := baseUrl.ResolveReference(p).String()
	res, err := http.Get(path)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	webAppRegex, err := regexp.Compile(WebAppRegex)
	if err != nil {
		panic(err)
	}

	return webAppRegex.FindStringSubmatch(string(body))[1]
}
