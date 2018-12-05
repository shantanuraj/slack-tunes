package provider

import (
	"github.com/shantanuraj/slack-tunes/logger"
)

const iTunes = "iTunes"

// ITunes fetches song info from iTunes
type ITunes struct {
	logger *logger.Logger
}

// CurrentSong fetches the current song from iTunes
func (i ITunes) CurrentSong() (Song, error) {
	song, err := currentSong(iTunes)

	if err != nil {
		i.logger.Log("[provider-itunes] Error in getting player state", err)
	}

	return song, err
}

// IsPlaying returns boolean to indicate if iTunes is playing
func (i ITunes) IsPlaying() (bool, error) {
	playing, err := isPlaying(iTunes)

	if err != nil {
		i.logger.Log("[provider-itunes] Error in getting player state", err)
	}

	return playing, err
}

// NewITunes returns an instance of the `ITunes` provider
func NewITunes() ITunes {
	return ITunes{
		logger: logger.GetLogger(),
	}
}
