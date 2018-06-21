package test

import (
	"testing"
	"time"

	. "github.com/Mungrel/steam-go"

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

	expectedLastLogOff := Timestamp(time.Unix(1529473982, 0))
	expectedTimeCreated := Timestamp(time.Unix(1063407589, 0))

	expected := []*PlayerSummary{
		&PlayerSummary{
			SteamID:                  "76561197960435530",
			CommunityVisibilityState: 3,
			ProfileState:             1,
			PersonaName:              "Robin",
			LastLogOff:               &expectedLastLogOff,
			ProfileURL:               "https://steamcommunity.com/id/robinwalker/",
			Avatar:                   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4.jpg",
			AvatarMedium:             "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4_medium.jpg",
			AvatarFull:               "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4_full.jpg",
			PersonaState:             0,
			RealName:                 "Robin Walker",
			PrimaryClanID:            "103582791429521412",
			TimeCreated:              &expectedTimeCreated,
			PersonaStateFlags:        0,
			CountryCode:              "US",
			StateCode:                "WA",
			CityID:                   3961,
		},
	}

	assert.Equal(expected, players)
}
