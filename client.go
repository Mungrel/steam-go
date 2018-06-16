package steam

// Client is a basic Steam API client
type Client struct {
	apiKey string
}

// NewSteamClient creates a new Client
func NewSteamClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}
