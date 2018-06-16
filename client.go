package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://api.steampowered.com"

// Client is a basic Steam API client
type Client struct {
	apiKey string
}

// NewSteamClient creates a new Client
func NewSteamClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (client *Client) get(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(respBytes, &result)
}
