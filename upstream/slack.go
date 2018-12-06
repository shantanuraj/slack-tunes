package upstream

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/shantanuraj/slack-tunes/logger"

	"github.com/shantanuraj/slack-tunes/provider"
)

// Slack updates the song info to a slack status
type Slack struct {
	api        *slack.Client
	logger     *logger.Logger
	lastStatus *string
}

func getStatusText(providerName string, isPlaying bool, s provider.Song) string {
	if isPlaying {
		return fmt.Sprintf("%s by %s on %s", s.Title, s.Artist, providerName)
	}
	return ""
}

func getStatusEmoji(isPlaying bool) string {
	if isPlaying {
		return ":headphones:"
	}
	return ""
}

// UpdateSong posts the song data to Slack's status API
func (s *Slack) UpdateSong(providerName string, isPlaying bool, song provider.Song) error {
	var err error
	status := getStatusText(providerName, isPlaying, song)
	emoji := getStatusEmoji(isPlaying)

	if s.lastStatus != nil && status == *s.lastStatus {
		s.logger.Log("[upstream-slack] Same status not updating api")
		return nil
	}
	s.lastStatus = &status

	if err = s.api.SetUserCustomStatus(status, emoji); err != nil {
		s.logger.Log("[upstream-slack] Could not update status", err)
	} else {
		if status == "" || emoji == "" {
			s.logger.Log("[upstream-slack] Resetted status")
		} else {
			s.logger.Log("[upstream-slack] Updated status to", emoji, status)
		}
	}

	return err
}

// NewSlack returns an instance of the slack upstream
func NewSlack() *Slack {
	apiToken := os.Getenv("STUNES_SLACK_TOKEN")
	if apiToken == "" {
		log.Fatal("Couldn't get environment variable for Slack")
	}
	return &Slack{
		api:        slack.New(apiToken),
		logger:     logger.GetLogger(),
		lastStatus: nil,
	}
}
