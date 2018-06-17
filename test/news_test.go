package test

import (
	"testing"

	"github.com/Mungrel/testify/assert"
)

const (
	appID     = "123456"
	count     = 0
	maxLength = 0
)

func TestNews(t *testing.T) {
	mockSteamClient := newMockClient()
	assert := assert.New(t)

	newsItems, err := mockSteamClient.GetNewsForApp(appID, count, maxLength)

	assert.Nil(err)
	assert.NotNil(newsItems)
	assert.NotEmpty(newsItems)
	assert.Len(newsItems, 3)
}
