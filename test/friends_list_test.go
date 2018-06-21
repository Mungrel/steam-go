package test

import (
	"testing"

	"github.com/Mungrel/testify/assert"
)

const steamID = "dummy-steam-id"
const relationship = "friend"

func TestFriendsList(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	friends, err := mockSteamClient.GetFriendsList(steamID, relationship)
	assert.Nil(err)
	assert.NotNil(friends)
	assert.NotEmpty(friends)
	assert.Len(friends, 9)

	for _, friend := range friends {
		assert.NotEmpty(friend.SteamID)
		assert.NotEmpty(friend.Relationship)
	}
}
