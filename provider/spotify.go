package provider

// Spotify app name
const Spotify = "Spotify"

// NewSpotify returns an instance of the `Spotify` provider
func NewSpotify() Provider {
	return NewMacProvider(Spotify)
}
