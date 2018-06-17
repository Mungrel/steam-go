package steam

import (
	"net/url"
	"strconv"
)

// NewsItem represents a NewsItem from the Steam API
type NewsItem struct {
	GID           *string    `json:"gid"`
	Title         *string    `json:"title"`
	URL           *string    `json:"url"`
	IsExternalURL bool       `json:"is_external_url"`
	Author        *string    `json:"author"`
	Contents      *string    `json:"contents"`
	FeedLabel     *string    `json:"feedlabel"`
	Date          *Timestamp `json:"date"`
	FeedName      *string    `json:"feedname"`
	FeedType      int64      `json:"feed_type"`
	AppID         int64      `json:"appid"`
}

type newsItemResponse struct {
	AppNews appNews `json:"appnews"`
}

type appNews struct {
	NewsItems []*NewsItem `json:"newsitems"`
}

// GetNewsForApp gets news items from the Steam API
func (client *Client) GetNewsForApp(appID string, count, maxLength int64) ([]*NewsItem, error) {
	parameters := url.Values{}
	parameters.Add("appid", appID)
	parameters.Add("count", strconv.FormatInt(count, 10))
	parameters.Add("maxlength", strconv.FormatInt(maxLength, 10))

	url := baseURL + "/ISteamNews/GetNewsForApp/v0002/?"
	url += parameters.Encode()

	var resp newsItemResponse
	err := client.get(url, &resp)
	if err != nil {
		return nil, err
	}

	return resp.AppNews.NewsItems, nil
}
