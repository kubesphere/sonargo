package response

import (
	"net/http"
	"net/url"
)

type Client struct {
	baseURL                   *url.URL
	username, password, token string
	authType                  authType
	httpClient                *http.Client
	Webservices               *WebservicesService
}

func NewClient(endpoint, username, password string) (*Client, error) {
	c := &Client{username: username, password: password, authType: basicAuth, httpClient: http.DefaultClient}
	if endpoint == "" {
		c.SetBaseURL(defaultBaseURL)
	} else {
		if err := c.SetBaseURL(endpoint); err != nil {
			return nil, err
		}
	}
	c.Webservices = &WebservicesService{client: c}
	return c, nil
}
