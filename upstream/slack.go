package upstream

import (
	"github.com/shantanuraj/slack-tunes/logger"

	"github.com/shantanuraj/slack-tunes/provider"
)

// Slack updates the song info to a slack status
type Slack struct {
	logger *logger.Logger
}

// UpdateSong posts the song data to Slack's status API
func (s Slack) UpdateSong(song provider.Song) error {
	s.logger.Log("[upstream-slack] Updating status to", song.Title, "by", song.Artist)
	return nil
}

// NewSlack returns an instance of the slack upstream
func NewSlack() Slack {
	return Slack{
		logger: logger.GetLogger(),
	}
}
