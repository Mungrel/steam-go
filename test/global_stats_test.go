package test

import (
	"testing"

	"github.com/Mungrel/testify/assert"
)

func TestGlobalStats(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	statNames := []string{"global.map.emp_isle"}

	stats, err := mockSteamClient.GetGlobalStatsForGame(appID, statNames)
	assert.Nil(err)
	assert.NotNil(stats)

	assert.Equal(stats.Result, 1)
	assert.Len(stats.Stats, 1)

	expected := map[string]string{
		"total": "28871095469",
	}
	assert.Equal(expected, stats.Stats)
}
