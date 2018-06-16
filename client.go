package steam

type SteamClient struct {
	apiKey string
}

func NewSteamClient(apiKey string) *SteamClient {
	return &SteamClient{
		apiKey: apiKey,
	}
}
