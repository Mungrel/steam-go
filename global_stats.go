package steam

import (
	"fmt"
	"net/url"
	"strconv"
)

type globalStatsResponse struct {
	Response GlobalStats `json:"response"`
}

// GlobalStats represents the stats from the Steam API
type GlobalStats struct {
	Stats  map[string]string `json:"globalstats"`
	Result int64             `json:"result"`
}

func (client *Client) GetGlobalStatsForGame(appID string, names []string) (*GlobalStats, error) {
	parameters := url.Values{}
	parameters.Add("appid", appID)
	parameters.Add("count", strconv.Itoa(len(names)))

	for i := range names {
		name := fmt.Sprintf("name[%d]", i)
		parameters.Add(name, names[i])
	}

	url := baseURL + "/ISteamUserStats/GetGlobalStatsForGame/v0001/"
	url += parameters.Encode()

	var response globalStatsResponse
	err := client.get(url, &response)
	if err != nil {
		return nil, err
	}

	return &response.Response, nil
}
