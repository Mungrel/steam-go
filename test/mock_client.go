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

	return nil, errors.New("Endpoint not mocked")
}

func newResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}
