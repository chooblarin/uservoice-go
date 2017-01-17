package uservoice

import (
	"errors"
	"net/http"
	"net/url"
	"path"

	"encoding/json"

	"github.com/mrjones/oauth"
)

// Client represents API client
type Client struct {
	URL      *url.URL
	consumer *oauth.Consumer
	token    *oauth.AccessToken
}

// NewClient returns new Client
func NewClient(subdomain, apiKey, apiSecret string) (*Client, error) {

	// validate
	if len(apiKey) == 0 {
		return nil, errors.New("missing api key")
	}
	if len(apiSecret) == 0 {
		return nil, errors.New("missing api secret")
	}
	urlString := "https://" + subdomain + ".uservoice.com"
	parsedURL, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("failed to parse url")
	}
	consumer := oauth.NewConsumer(apiKey, apiSecret, oauth.ServiceProvider{})
	consumer.Debug(true) // TODO: disable debug flag
	client := &Client{
		URL:      parsedURL,
		consumer: consumer,
	}
	return client, nil
}

// Request returns API response
func (c *Client) Request(method, spath string, params map[string]string) (*http.Response, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	url := u.String()

	switch method {
	case "GET":
		return c.consumer.Get(url, nil, &oauth.AccessToken{})
	case "POST":
		m, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		return c.consumer.PostJson(url, string(m), &oauth.AccessToken{})
	default:
		// TODO: Support PUT and DELETE
		return nil, errors.New("PUT and DELETE calls are not supported")
	}
}
