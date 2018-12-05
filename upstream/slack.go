package upstream

import (
	"github.com/shantanuraj/slack-tunes/logger"

	"github.com/shantanuraj/slack-tunes/provider"
)

// Slack updates the song info to a slack status
type Slack struct{}

// UpdateSong posts the song data to Slack's status API
func (i Slack) UpdateSong(song provider.Song) error {
	l := logger.GetLogger()
	l.Log("[upstream-slack] Updating status to", song.Title, "by", song.Artist)
	return nil
}
