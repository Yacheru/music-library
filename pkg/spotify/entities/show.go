package entities

type SimpleShowPage struct {
	basePage
	Shows []FullShow `json:"items"`
}

type SimpleShow struct {
	// A list of the countries in which the show can be played,
	// identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`

	// The copyright statements of the show.
	Copyrights []Copyright `json:"copyrights"`

	// A description of the show.
	Description string `json:"description"`

	// Whether or not the show has explicit content
	// (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`

	// Known external URLs for this show.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details
	// of the show.
	Href string `json:"href"`

	// The SpotifyID for the show.
	ID ID `json:"id"`

	// The cover art for the show in various sizes,
	// widest first.
	Images []Image `json:"images"`

	// True if all of the show’s episodes are hosted outside
	// of Spotify’s CDN. This field might be null in some cases.
	IsExternallyHosted *bool `json:"is_externally_hosted"`

	// A list of the languages used in the show, identified by
	// their ISO 639 code.
	Languages []string `json:"languages"`

	// The media type of the show.
	MediaType string `json:"media_type"`

	// The name of the show.
	Name string `json:"name"`

	// The publisher of the show.
	Publisher string `json:"publisher"`

	// The object type: “show”.
	Type string `json:"type"`

	// The Spotify URI for the show.
	URI URI `json:"uri"`
}

type ResumePointObject struct {
	// 	Whether or not the episode has been fully played by the user.
	FullyPlayed bool `json:"fully_played"`

	// The user’s most recent position in the episode in milliseconds.
	ResumePositionMs Numeric `json:"resume_position_ms"`
}

type FullShow struct {
	SimpleShow

	// A list of the show’s episodes.
	Episodes SimpleEpisodePage `json:"episodes"`
}
