package steam

import (
	"net/url"
)

type friendsListResponse struct {
	FriendsList friendsList `json:"friendsList"`
}

type friendsList struct {
	Friends []*Friend `json:"friends"`
}

// Friend represents a friend in the Steam API
type Friend struct {
	SteamID      string     `json:"steamid"`
	Relationship string     `json:"relationship"`
	FriendSince  *Timestamp `json:"friend_since"`
}

// GetFriendsList gets a list of friends from the Steam API
func (client *Client) GetFriendsList(steamID, relationship string) ([]*Friend, error) {
	parameters := url.Values{}
	parameters.Add("key", client.APIKey)
	parameters.Add("steamid", steamID)
	parameters.Add("relationship", relationship)

	url := baseURL + "/ISteamUser/GetFriendList/v0001/"
	url += parameters.Encode()

	var response friendsListResponse
	err := client.get(url, &response)
	if err != nil {
		return nil, err
	}

	return response.FriendsList.Friends, nil
}
