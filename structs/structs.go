package structs

type StanRequest struct {
	Payload []struct {
		Country string`json:"country"`
		Description string`json:"description"`
		Drm bool `json:"drm"`
		EpisodeCount int `json:"episodeCount"`
		Genre string `json:"genre"`
		Image struct {
			ShowImage string `json:"showImage"`
		} `json:"image"`
		Language string `json:"language"`
		NextEpisode struct {
			Channel string `json:"channel"`
			ChannelLogo string `json:"channelLogo"`
			Date string `json:"date"`
			Html string `json:"html"`
			Url string `json:"url"`
		} `json:"nextEpisode"`
		PrimaryColour string `json:"primaryColour"`
		Seasons []struct {
			Slug string `json:"slug"`
		} `json:"seasons"`
		Slug string `json:"slug"`
		Title string `json:"title"`
		TvChannel string `json:"tvChannel"`
	} `json:"payload"`
	Skip int `json:"skip"`
	Take int `json:"take"`
	TotalRecords int `json:"totalRecords"`
}

type StanResponse struct {
	Response []ResponseMovie `json:"response"`
}

type ResponseMovie struct {
	Image string`json:"image"`
	Slug string`json:"slug"`
	Title string `json:"title"`
}