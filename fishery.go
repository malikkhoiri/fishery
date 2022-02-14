package fishery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func (fc *Client) get(sheetName string, s, t interface{}) (err error) {
	search, err := json.Marshal(s)

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/%s/%s?search=%s", fc.baseUrl, fc.apiKey, sheetName, string(search))
	err = fc.call(http.MethodGet, url, nil, t)

	return
}

func (fc *Client) add(sheetName, data, t interface{}) (err error) {
	body, err := json.Marshal(data)

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/%s/%s", fc.baseUrl, fc.apiKey, sheetName)
	err = fc.call(http.MethodPost, url, bytes.NewReader(body), t)

	return
}

func (fc *Client) update(sheetName, data, t interface{}) (err error) {
	body, err := json.Marshal(data)

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/%s/%s", fc.baseUrl, fc.apiKey, sheetName)
	err = fc.call(http.MethodPut, url, bytes.NewReader(body), t)

	return
}

func (fc *Client) delete(sheetName, data, t interface{}) (err error) {
	body, err := json.Marshal(data)

	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/%s/%s", fc.baseUrl, fc.apiKey, sheetName)
	err = fc.call(http.MethodDelete, url, bytes.NewReader(body), t)

	return
}

func (fc *Client) call(method, url string, body io.Reader, response interface{}) (err error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return
	}

	res, err := fc.Client.Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	if response == nil {
		return
	}

	err = json.Unmarshal(resBody, response)

	if err != nil {
		return
	}

	return
}
