package test

import (
	"testing"

	. "github.com/Mungrel/steam-go"

	"github.com/Mungrel/testify/assert"
)

func TestGameSchema(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	schema, err := mockSteamClient.GetSchemaForGame(appID, nil)
	assert.Nil(err)
	assert.NotNil(schema)

	expected := &GameSchema{
		Name:    "PAYDAY 2",
		Version: "538",
		AvailableGameStats: &GameStats{
			Achievements: []*Achievement{
				{
					Name:         "hot_wheels",
					DefaultValue: 0,
					DisplayName:  "Coming in Hot",
					Hidden:       0,
					Description:  "On day 1 of the Watchdogs job, don't let the cops shoot and destroy the escape vehicle.",
					IconURL:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/218620/f6ed9cd6ec9750bcd36193c74e6f16104f6c1267.jpg",
					IconGrayURL:  "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/218620/c336adacd88a21a6010c9b5596192322aecaf265.jpg",
				},
				{
					Name:         "fish_ai",
					DefaultValue: 0,
					DisplayName:  "Fish A.I.",
					Hidden:       0,
					Description:  "On day 2 of the Watchdogs job, throw a loot bag into the sea, hoping fish move away as it gets near.",
					IconURL:      "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/218620/19fd9f952f27bad6bc909721a69c6374cfa5e31b.jpg",
					IconGrayURL:  "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/218620/6f6357e4e4106e7b120c2ef6814061e1c6e989ec.jpg",
				},
			},
			Stats: []*Stat{
				{
					Name:         "player_time",
					DefaultValue: 0,
					DisplayName:  "",
				},
				{
					Name:         "player_time_0h",
					DefaultValue: 0,
					DisplayName:  "",
				},
			},
		},
	}

	assert.Equal(expected, schema)
}
