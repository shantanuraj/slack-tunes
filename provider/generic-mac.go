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

// GetName returns the name of the provider
func (g GenericMacProvider) GetName() string {
	return g.appName
}

// CurrentSong fetches the current song from app
func (g GenericMacProvider) CurrentSong() (Song, error) {
	song, err := currentSong(g.appName)

	if err != nil {
		g.logger.Log(g.debugPrefix, "Error in getting player state", err)
	}

	return song, err
}

// IsPlaying returns boolean to indicate if app is playing
func (g GenericMacProvider) IsPlaying() (bool, error) {
	playing, err := isPlaying(g.appName)

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
