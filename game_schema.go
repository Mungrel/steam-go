package steam

import (
	"net/url"
)

type schemaResponse struct {
	GameSchema GameSchema `json:"game"`
}

// GameSchema represents the game schema object in the Steam API response
type GameSchema struct {
	Name               string     `json:"gameName"`
	Version            string     `json:"gameVersion"`
	AvailableGameStats *GameStats `json:"availableGameStats"`
}

// GameStats represents the game stats object in the Steam API response
type GameStats struct {
	Achievements []*Achievement `json:"achievements"`
	Stats        []*Stat        `json:"stats"`
}

// Achievement represents an achievement from the Steam API game schema
type Achievement struct {
	Name         string `json:"name"`
	DefaultValue int64  `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
	Hidden       int    `json:"hidden"`
	Description  string `json:"description"`
	IconURL      string `json:"icon"`
	IconGrayURL  string `json:"icongray"`
}

// Stat represents a stat in the Steam API game schema
type Stat struct {
	Name         string `json:"name"`
	DefaultValue int64  `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
}

// GetSchemaForGame gets the game schema for an app in a specified language.
// Language will default to English if not specified
func (client *Client) GetSchemaForGame(appID string, language *string) (*GameSchema, error) {
	parameters := url.Values{}
	parameters.Add("key", client.APIKey)
	parameters.Add("appid", appID)
	if language != nil {
		parameters.Add("l", *language)
	}

	url := baseURL + "/ISteamUserStats/GetSchemaForGame/v2/"
	url += parameters.Encode()

	var response schemaResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return &response.GameSchema, nil
}
