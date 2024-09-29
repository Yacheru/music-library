package entities

type FullTrackPage struct {
	basePage
	Tracks []FullTrack `json:"items"`
}

type SimpleTrack struct {
	Album   SimpleAlbum    `json:"album"`
	Artists []SimpleArtist `json:"artists"`
	// A list of the countries in which the track can be played,
	// identified by their ISO 3166-1 alpha-2 codes.
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually 1 unless the album consists of more than one disc).
	DiscNumber Numeric `json:"disc_number"`
	// The length of the track, in milliseconds.
	Duration Numeric `json:"duration_ms"`
	// Whether or not the track has explicit lyrics.
	// true => yes, it does; false => no, it does not.
	Explicit bool `json:"explicit"`
	// External URLs for this track.
	ExternalURLs map[string]string `json:"external_urls"`
	// ExternalIDs are IDs for this track in other databases
	ExternalIDs TrackExternalIDs `json:"external_ids"`
	// A link to the Web API endpoint providing full details for this track.
	Endpoint string `json:"href"`
	ID       ID     `json:"id"`
	Name     string `json:"name"`
	// A URL to a 30 second preview (MP3) of the track.
	PreviewURL string `json:"preview_url"`
	// The number of the track.  If an album has several
	// discs, the track number is the number on the specified
	// DiscNumber.
	TrackNumber Numeric `json:"track_number"`
	URI         URI     `json:"uri"`
	// Type of the track
	Type string `json:"type"`
}

type FullTrack struct {
	SimpleTrack
	// The album on which the track appears. The album object includes a link in href to full information about the album.
	Album SimpleAlbum `json:"album"`
	// Known external IDs for the track.
	ExternalIDs map[string]string `json:"external_ids"`
	// Popularity of the track.  The value will be between 0 and 100,
	// with 100 being the most popular.  The popularity is calculated from
	// both total plays and most recent plays.
	Popularity Numeric `json:"popularity"`

	// IsPlayable defines if the track is playable. It's reported when the "market" parameter is passed to the tracks
	// listing API.
	// See: https://developer.spotify.com/documentation/general/guides/track-relinking-guide/
	IsPlayable *bool `json:"is_playable"`

	// LinkedFrom points to the linked track. It's reported when the "market" parameter is passed to the tracks listing
	// API.
	LinkedFrom *LinkedFromInfo `json:"linked_from"`
}

type LinkedFromInfo struct {
	// ExternalURLs are the known external APIs for this track or album
	ExternalURLs map[string]string `json:"external_urls"`

	// Href is a link to the Web API endpoint providing full details
	Href string `json:"href"`

	// ID of the linked track
	ID ID `json:"id"`

	// Type of the link: album of the track
	Type string `json:"type"`

	// URI is the Spotify URI of the track/album
	URI string `json:"uri"`
}
