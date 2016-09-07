package trakitapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client

	ApiUrl    *url.URL
	WebAppUrl *url.URL
	AuthToken string

	Customers    CustomerService
	Landmarks    LandmarksService
	SessionToken SessionTokenService
	Users        UserService
	Versions     VersionsService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient}
	c.Customers = &CustomerServiceOp{client: c}
	c.Landmarks = &LandmarkServiceOp{client: c}
	c.SessionToken = &SessionTokenServiceOp{client: c}
	c.Users = &UserServiceOp{client: c}
	c.Versions = &VersionsServiceOp{client: c}

	return c
}

func (c *Client) SetApiUrl(in string) *Client {
	c.ApiUrl, _ = url.Parse(in)
	return c
}

func (c *Client) SetWebAppUrl(in string) *Client {
	c.WebAppUrl, _ = url.Parse(in)
	return c
}

func (c *Client) SetAuthToken(in string) *Client {
	c.AuthToken = in
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.ApiUrl.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if c.AuthToken != "" {
		req.Header.Add("X-Auth-Token", c.AuthToken)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err := json.NewDecoder(resp.Body).Decode(&v)
			if err != nil {
				return resp, err
			}
		}
	}

	return resp, err
}
