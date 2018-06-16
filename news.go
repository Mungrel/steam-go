package steam

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
	FeedType      *string    `json:"feed_type"`
	AppID         *string    `json:"appid"`
}
