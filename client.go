package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://api.steampowered.com"

type clientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a basic Steam API client
type Client struct {
	APIKey string
	Client clientInterface
}

// NewSteamClient creates a new Client
func NewSteamClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Client: http.DefaultClient,
	}
}

func (c *Client) get(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(respBytes, &result)
}
