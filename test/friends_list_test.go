package test

import (
	"testing"
	"time"

	. "github.com/Mungrel/steam-go"
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
	assert.Len(friends, 2)

	expectedFriendSince := Timestamp(time.Unix(0, 0))
	expected := []*Friend{
		&Friend{
			SteamID:      "76561197960265731",
			Relationship: "friend",
			FriendSince:  &expectedFriendSince,
		},
		&Friend{
			SteamID:      "76561197960265738",
			Relationship: "friend",
			FriendSince:  &expectedFriendSince,
		},
	}

	assert.Equal(expected, friends)
}
