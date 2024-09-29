package entities

type Image struct {
	// The image height, in pixels.
	Height Numeric `json:"height"`
	// The image width, in pixels.
	Width Numeric `json:"width"`
	// The source URL of the image.
	URL string `json:"url"`
}
