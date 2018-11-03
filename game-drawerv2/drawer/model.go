package drawer

// RawResponse form statsapi
type RawResponse struct {
	LiveData struct {
		Plays struct {
			AllPlays []struct {
				Result struct {
					EventTypeID string `json:"eventTypeId"`
				} `json:"result"`
				Coordinates struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"coordinates"`
			} `json:"allPlays"`
		} `json:"plays"`
	} `json:"liveData"`
}
