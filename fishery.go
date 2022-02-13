package fishery

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	*http.Client
	baseUrl string
	apiKey  string
}

type Search map[string]interface{}

func NewClient(baseUrl, apiKey string) (*Client, error) {
	url, err := urlParse(baseUrl)

	if err != nil {
		return nil, err
	}

	return &Client{
		Client:  http.DefaultClient,
		baseUrl: url,
		apiKey:  apiKey,
	}, nil
}

func (c *Client) GetSheet(sheetName string) *Sheet {
	return &Sheet{
		client: c,
		name:   sheetName,
	}
}

func urlParse(baseUrl string) (string, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	if url.Scheme == "" {
		return "", fmt.Errorf("http or https scheme not specified")
	}

	if url.Scheme != "https" && url.Scheme != "http" {
		return "", fmt.Errorf("scheme must be use http or https")
	}

	return url.String(), nil
}
