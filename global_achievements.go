package steam

import (
	"net/url"
)

type globalAchievementsResponse struct {
	AchievementPercentages achievementPercentages `json:"achievementpercentages"`
}

type achievementPercentages struct {
	Achievements []*AchievementPercentage `json:"achievements"`
}

// AchievementPercentage represents an acheivement and its global attainment percentage
type AchievementPercentage struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"percent"`
}

// GetGlobalAchievementPercentagesForApp gets global achievement percentages for an app from the Steam API
func (client *Client) GetGlobalAchievementPercentagesForApp(appID string) ([]*AchievementPercentage, error) {
	parameters := url.Values{}
	parameters.Add("gameid", appID)

	url := baseURL + "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/"
	url += parameters.Encode()

	var response globalAchievementsResponse
	err := client.get(url, &response)
	if err != nil {
		return nil, err
	}

	return response.AchievementPercentages.Achievements, nil
}
