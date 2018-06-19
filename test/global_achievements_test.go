package test

import (
	"testing"

	"github.com/Mungrel/testify/assert"
)

func TestGlobalAchievements(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	achievements, err := mockSteamClient.GetGlobalAchievementPercentagesForApp(appID)
	assert.Nil(err)
	assert.NotNil(achievements)
	assert.NotEmpty(achievements)
	assert.Len(achievements, 1)
}
