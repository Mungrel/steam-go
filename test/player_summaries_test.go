package test

import (
	"testing"

	"github.com/Mungrel/testify/assert"
)

func TestPlayerSummaries(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	players, err := mockSteamClient.GetPlayerSummaries([]string{steamID})
	assert.Nil(err)
	assert.NotNil(players)
	assert.NotEmpty(players)
	assert.Len(players, 1)

	for _, friend := range players {
		assert.NotEmpty(friend.SteamID)
	}
}
