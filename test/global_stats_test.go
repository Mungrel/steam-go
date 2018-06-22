package test

import (
	"testing"

	. "github.com/Mungrel/steam-go"
	"github.com/Mungrel/testify/assert"
)

func TestGlobalStats(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	statNames := []string{"global.map.emp_isle"}

	stats, err := mockSteamClient.GetGlobalStatsForGame(appID, statNames)
	assert.Nil(err)
	assert.NotNil(stats)

	assert.Equal(stats.Result, int64(1))
	assert.Len(stats.Stats, 1)

	expectedStat := &GlobalStat{
		Total: "28871095469",
	}
	expectedStats := map[string]*GlobalStat{
		"global.map.emp_isle": expectedStat,
	}
	assert.Equal(expectedStats, stats.Stats)
}
