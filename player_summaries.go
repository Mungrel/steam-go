package steam

import (
	"net/url"
	"strings"
)

type playerSummariesResponse struct {
	Response playerSummaries `json:"response"`
}

type playerSummaries struct {
	Players []*PlayerSummary `json:"players"`
}

// PlayerSummary represents a player summary from the Steam API
type PlayerSummary struct {
	SteamID                  string     `json:"steamid"`
	CommunityVisibilityState int        `json:"communityvisibilitystate"`
	ProfileState             int        `json:"profilestate"`
	PersonaName              string     `json:"personaname"`
	LastLogOff               *Timestamp `json:"lastlogoff"`
	ProfileURL               string     `json:"profileurl"`
	Avatar                   string     `json:"avatar"`
	AvatarMedium             string     `json:"avatarmedium"`
	AvatarFull               string     `json:"avatarfull"`
	PersonaState             int        `json:"personastate"`
	RealName                 string     `json:"realname"`
	PrimaryClanID            string     `json:"primaryclanid"`
	TimeCreated              *Timestamp `json:"timecreated"`
	PersonaStateFlags        int        `json:"personastateflags"`
	CountryCode              string     `json:"loccountrycode"`
	StateCode                string     `json:"locstatecode"`
	CityID                   int64      `json:"loccityid"`
}

// GetPlayerSummaries retrieves player summaries from the Steam API given those players' IDs
func (client *Client) GetPlayerSummaries(steamIDs []string) ([]*PlayerSummary, error) {
	parameters := url.Values{}
	parameters.Add("key", client.APIKey)
	parameters.Add("steamids", strings.Join(steamIDs, ","))

	url := baseURL + "/ISteamUser/GetPlayerSummaries/v0002/"
	url += parameters.Encode()

	var response playerSummariesResponse
	err := client.get(url, &response)
	if err != nil {
		return nil, err
	}

	return response.Response.Players, nil
}
