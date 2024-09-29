package entities

type SearchResult struct {
	Artists   *FullArtistPage     `json:"artists"`
	Albums    *SimpleAlbumPage    `json:"albums"`
	Playlists *SimplePlaylistPage `json:"playlists"`
	Tracks    *FullTrackPage      `json:"tracks"`
	Shows     *SimpleShowPage     `json:"shows"`
	Episodes  *SimpleEpisodePage  `json:"episodes"`
}
