package provider

import (
	"strings"

	"github.com/shantanuraj/slack-tunes/logger"
)

// GenericMacProvider defines a generic mac based music app
type GenericMacProvider struct {
	appName     string
	debugPrefix string
	logger      *logger.Logger
}

// CurrentSong fetches the current song from iTunes
func (g GenericMacProvider) CurrentSong() (Song, error) {
	song, err := currentSong(iTunes)

	if err != nil {
		g.logger.Log(g.debugPrefix, "Error in getting player state", err)
	}

	return song, err
}

// IsPlaying returns boolean to indicate if iTunes is playing
func (g GenericMacProvider) IsPlaying() (bool, error) {
	playing, err := isPlaying(iTunes)

	if err != nil {
		g.logger.Log(g.debugPrefix, "Error in getting player state", err)
	}

	return playing, err
}

// NewMacProvider returns an instance of `Provider`
func NewMacProvider(appName string) Provider {
	return GenericMacProvider{
		appName:     appName,
		debugPrefix: "[provider-" + strings.ToLower(appName) + "]",
		logger:      logger.GetLogger(),
	}
}
