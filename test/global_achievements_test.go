package test

import (
	"testing"

	. "github.com/Mungrel/steam-go"
	"github.com/Mungrel/testify/assert"
)

func TestGlobalAchievements(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	achievements, err := mockSteamClient.GetGlobalAchievementPercentagesForApp(appID)
	assert.Nil(err)
	assert.NotNil(achievements)
	assert.NotEmpty(achievements)
	assert.Len(achievements, 2)

	expected := []*AchievementPercentage{
		{
			Name:       "TF_SCOUT_LONG_DISTANCE_RUNNER",
			Percentage: 53.20844650268555,
		},
		{
			Name:       "TF_HEAVY_DAMAGE_TAKEN",
			Percentage: 44.763389587402344,
		},
	}

	assert.Equal(expected, achievements)
}
