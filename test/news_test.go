package test

import (
	"testing"

	. "github.com/Mungrel/steam-go"

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
	assert.Len(newsItems, 2)

	expected := []*NewsItem{
		{
			GID:           "2431262783620979755",
			Title:         "Top Ten Plays from the Copenhagen LAN 2018",
			URL:           "https://steamstore-a.akamaihd.net/news/externalpost/tf2_blog/2431262783620979755",
			IsExternalURL: true,
			Author:        "",
			Contents:      "<a href=\"https://www.youtube.com/user/LuckyLuke0316\" target=\"_blank\">LuckyLukeTF2</a> is back with the <a href=\"https://www.youtube.com/watch?v=gPmgD3Bm9Lg\" target=\"_blank\">Top 10 Plays from the Copenhagen LAN 2018</a>! The Top 10 Plays is a monthly series and you can catch up on all of the episodes <a href=\"https://www.youtube.com/playlist?list=PLBD72F08BA43FC9C0\" target=\"_blank\">here</a>. <a href=\"https://www.youtube.com/user/DanishActionHero\" target=\"_blank\">DanishActionHero</a> has also created an exciting video of the <a href=\"https://www.youtube.com/watch?v=HMiTqxf_uC4\" target=\"_blank\">best frags and moments</a>! ",
			FeedLabel:     "TF2 Blog",
			Date:          NewTimestamp(1528907340),
			FeedName:      "tf2_blog",
			FeedType:      0,
			AppID:         440,
		},
		{
			GID:           "2431262152935892422",
			Title:         "Team Fortress 2 Update Released",
			URL:           "https://steamstore-a.akamaihd.net/news/externalpost/tf2_blog/2431262152935892422",
			IsExternalURL: true,
			Author:        "",
			Contents:      "An update to Team Fortress 2 has been released. The update will be applied automatically when you restart Team Fortress 2. The major changes include: * Several fixes to address exploits and stability issues ;* Fixed server crash related to the CTriggerAddTFPlayerCondition map entity ;* Fixed client ...",
			FeedLabel:     "TF2 Blog",
			Date:          NewTimestamp(1528402980),
			FeedName:      "tf2_blog",
			FeedType:      0,
			AppID:         440,
		},
	}

	assert.Equal(expected, newsItems)
}
