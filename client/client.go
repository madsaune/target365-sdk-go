package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	BaseURLShared     = "https://shared.target365.io/api"
	BaseURLTest       = "https://test.target365.io/api"
	headerAccept      = "application/json"
	headerContentType = "application/json"
	headerAPIKey      = "X-ApiKey"
)

type Client struct {
	BaseURL  string
	APIToken string
	client   *http.Client
}

func NewClient(baseURL string, token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		BaseURL:  baseURL,
		APIToken: token,
		client:   httpClient,
	}

	return c
}

func (c *Client) Get(path string, queryParams map[string]string, headers map[string]string) (*http.Response, error) {
	URL, err := url.Parse(fmt.Sprintf("%s%s", c.BaseURL, path))
	if err != nil {
		return nil, err
	}

	var params url.Values
	for k, v := range queryParams {
		params.Add(k, v)
	}

	URL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, URL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAPIKey, c.APIToken)
	req.Header.Add("Accept", headerAccept)
	req.Header.Add("Content-Type", headerContentType)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Post(path string, body []byte, headers map[string]string) (*http.Response, error) {
	URL := fmt.Sprintf("%s%s", c.BaseURL, path)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAPIKey, c.APIToken)
	req.Header.Add("Accept", headerAccept)
	req.Header.Add("Content-Type", headerContentType)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Delete(path string) (*http.Response, error) {
	URL := fmt.Sprintf("%s%s", c.BaseURL, path)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, URL, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAPIKey, c.APIToken)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
