package entities

type SimpleEpisodePage struct {
	basePage
	Episodes []EpisodePage `json:"items"`
}

type EpisodePage struct {
	// A URL to a 30 second preview (MP3 format) of the episode.
	AudioPreviewURL string `json:"audio_preview_url"`

	// A description of the episode.
	Description string `json:"description"`

	// The episode length in milliseconds.
	Duration_ms Numeric `json:"duration_ms"`

	// Whether or not the episode has explicit content
	// (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`

	// 	External URLs for this episode.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details of the episode.
	Href string `json:"href"`

	// The Spotify ID for the episode.
	ID ID `json:"id"`

	// The cover art for the episode in various sizes, widest first.
	Images []Image `json:"images"`

	// True if the episode is hosted outside of Spotify’s CDN.
	IsExternallyHosted bool `json:"is_externally_hosted"`

	// True if the episode is playable in the given market.
	// Otherwise false.
	IsPlayable bool `json:"is_playable"`

	// A list of the languages used in the episode, identified by their ISO 639 code.
	Languages []string `json:"languages"`

	// The name of the episode.
	Name string `json:"name"`

	// The date the episode was first released, for example
	// "1981-12-15". Depending on the precision, it might
	// be shown as "1981" or "1981-12".
	ReleaseDate string `json:"release_date"`

	// The precision with which release_date value is known:
	// "year", "month", or "day".
	ReleaseDatePrecision string `json:"release_date_precision"`

	// The user’s most recent position in the episode. Set if the
	// supplied access token is a user token and has the scope
	// user-read-playback-position.
	ResumePoint ResumePointObject `json:"resume_point"`

	// The show on which the episode belongs.
	Show SimpleShow `json:"show"`

	// The object type: "episode".
	Type string `json:"type"`

	// The Spotify URI for the episode.
	URI URI `json:"uri"`
}
