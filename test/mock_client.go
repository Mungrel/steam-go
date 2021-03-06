package test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	steam "github.com/Mungrel/steam-go"
)

type mockClient struct{}

func newMockClient() *steam.Client {
	return &steam.Client{
		APIKey: "dummy-key",
		Client: &mockClient{},
	}
}

func (mc *mockClient) Do(req *http.Request) (*http.Response, error) {
	url := req.URL.Path

	if strings.Contains(url, "/ISteamNews/GetNewsForApp/v0002/") {
		return newResponse(http.StatusOK, getJSON(news)), nil
	}

	if strings.Contains(url, "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/") {
		return newResponse(http.StatusOK, getJSON(globalAchievements)), nil
	}

	if strings.Contains(url, "/ISteamUserStats/GetGlobalStatsForGame/v0001/") {
		return newResponse(http.StatusOK, getJSON(globalStats)), nil
	}

	if strings.Contains(url, "/ISteamUser/GetPlayerSummaries/v0002/") {
		return newResponse(http.StatusOK, getJSON(playerSummaries)), nil
	}

	if strings.Contains(url, "/ISteamUser/GetFriendList/v0001/") {
		return newResponse(http.StatusOK, getJSON(friendsList)), nil
	}

	if strings.Contains(url, "/ISteamUserStats/GetSchemaForGame/v2/") {
		return newResponse(http.StatusOK, getJSON(gameSchema)), nil
	}

	return nil, errors.New("Endpoint not mocked")
}

func newResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}
