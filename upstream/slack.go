package upstream

import (
	"fmt"
	"log"
	"os"

	"github.com/shantanuraj/slack-tunes/logger"

	"github.com/slack-go/slack"

	"github.com/shantanuraj/slack-tunes/provider"
)

// Slack updates the song info to a slack status
type Slack struct {
	api        *slack.Client
	logger     *logger.Logger
	lastStatus *string
}

const slackMaxStatusLength = 100

func getStatusText(providerName string, isPlaying bool, s provider.Song) string {
	if isPlaying {
		status := fmt.Sprintf("%s by %s on %s", s.Title, s.Artist, providerName)
		if len(status) > slackMaxStatusLength {
			return status[:slackMaxStatusLength-3] + "..."
		}
		return status
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

	auth, err := s.api.AuthTest()
	if err != nil {
		s.logger.Log("[upstream-slack]", "Could not get auth info", err)
	}

	if err = s.api.SetUserCustomStatusWithUser(auth.UserID, status, emoji, 0); err != nil {
		s.logger.Log("[upstream-slack] Could not update status", err, 0)
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
