package thing

type RawResponse struct {
	Highlights struct {
		Scoreboard Highlights `json:"scoreboard"`
		GameCenter Highlights `json:"gameCenter"`
	} `json:"highlights"`
}

type Highlights struct {
	// Title     string `json:"title"`
	// TopicList string `json:"topicList"`
	Items []struct {
		// Type            string `json:"type"`
		// ID              string `json:"id"`
		// Date            string `json:"date"`
		// Title           string `json:"title"`
		// Blurb           string `json:"blurb"`
		Description string `json:"description"`
		// Duration        string `json:"duration"`
		// AuthFlow        bool   `json:"authFlow"`
		// MediaPlaybackID string `json:"mediaPlaybackId"`
		// MediaState      string `json:"mediaState"`
		// Keywords        []struct {
		// 	Type        string `json:"type"`
		// 	Value       string `json:"value"`
		// 	DisplayName string `json:"displayName"`
		// } `json:"keywords"`
		// Playbacks []struct {
		// 	Name   string `json:"name"`
		// 	Width  string `json:"width"`
		// 	Height string `json:"height"`
		// 	URL    string `json:"url"`
		// } `json:"playbacks"`
	} `json:"items"`
}
