package entities

type Followers struct {
	// The total number of followers.
	Count Numeric `json:"total"`
	// A link to the Web API endpoint providing full details of the followers,
	// or the empty string if this data is not available.
	Endpoint string `json:"href"`
}
