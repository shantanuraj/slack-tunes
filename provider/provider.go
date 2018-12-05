package provider

// Song represents a song playing in any provider
type Song struct {
	Title  string
	Artist string
}

// Provider is a generic type representing a song provider
type Provider interface {
	// CurrentSong returns a currently playing song
	CurrentSong() (Song, error)
}

// GetProvider returns the actual provivder for a given provider name
func GetProvider(providerName string) Provider {
	var p Provider
	// Only supporting iTunes for now
	p = ITunes{}
	return p
}
