package outmessage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/madsaune/target365-sdk-go/client"
)

type Client struct {
	BaseClient *client.Client
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseClient: client.NewClient(baseURL, token, nil),
	}
}

// Gets an out-message
func (c *Client) Get(transactionID string) (*Message, error) {
	if transactionID == "" {
		return nil, fmt.Errorf("transactionID cannot be empty")
	}

	URL := fmt.Sprintf("/out-messages/%s", transactionID)
	res, err := c.BaseClient.Get(URL, nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var m Message
	err = json.Unmarshal(raw, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// Posts a new out-message
func (c *Client) Create(om *Message) (*Response, error) {
	if err := om.Validate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(&om)
	if err != nil {
		return nil, err
	}

	log.Println(string(data))

	res, err := c.BaseClient.Post("/out-messages", data, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 201:
		loc := res.Header.Get("Location")
		id := strings.Split(loc, "/")

		return &Response{
			TransactionID: id[len(id)-1],
			Location:      res.Header.Get("Location"),
		}, nil
	case 400:
		return nil, fmt.Errorf("request had invalid payload")
	case 401:
		return nil, fmt.Errorf("request was unauthorized")
	default:
		return nil, fmt.Errorf("expected 201, got %d", res.StatusCode)
	}
}

// Posts a new batch of up to 100 out-messages
func (c *Client) CreateBatch(om []Message) ([]Response, error) {
	var responses []Response

	for _, v := range om {
		if err := v.Validate(); err != nil {
			return nil, err
		}

		responses = append(
			responses,
			Response{
				TransactionID: *v.TransactionID,
				Location:      fmt.Sprintf("%s%s%s", c.BaseClient.BaseURL, "/out-messages/", *v.TransactionID),
			},
		)
	}

	data, err := json.Marshal(om)
	if err != nil {
		return nil, err
	}

	res, err := c.BaseClient.Post("/out-messages/batch", data, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 201:
		return responses, nil
	case 400:
		return nil, fmt.Errorf("request had invalid payload")
	case 401:
		return nil, fmt.Errorf("request was unauthorized")
	default:
		return nil, fmt.Errorf("expected 201, got %d", res.StatusCode)
	}
}

// Deletes a future scheduled out-message
func (c *Client) Delete(transactionID string) error {
	if transactionID == "" {
		return fmt.Errorf("transactionID cannot be empty")
	}

	URL := fmt.Sprintf("/out-messages/%s", transactionID)
	res, err := c.BaseClient.Delete(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 204:
		return nil
	case 404:
		return fmt.Errorf("out-message not found")
	case 409:
		return fmt.Errorf("out-message could not be deleted")
	default:
		return fmt.Errorf("expected 201, got %d", res.StatusCode)
	}
}
