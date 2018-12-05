package upstream

import "github.com/shantanuraj/slack-tunes/provider"

// Upstream represents a generic API that accepts a Song and updates the api
type Upstream interface {
	// UpdateSong sends the song info to the Upstreams' status API
	UpdateSong(song provider.Song) error
}

// GetUpstream returns the actual upstream for a given upstream name
func GetUpstream(providerName string) Upstream {
	var u Upstream
	// Only supporting Slack for now
	u = Slack{}
	return u
}
