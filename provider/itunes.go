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
	song := Song{
		Title:  "I Am",
		Artist: "Jorja Smith",
	}
	return song, nil
}

// IsPlaying returns boolean to indicate if iTunes is playing
func (i ITunes) IsPlaying() (bool, error) {
	var out string
	var err error
	if out, err = tell(iTunes, playerStateCmd); err != nil {
		i.logger.Log("[provider-itunes] Error in getting player state", err)
	}
	return out == "playing", err
}

// NewITunes returns an instance of the `ITunes` provider
func NewITunes() ITunes {
	return ITunes{
		logger: logger.GetLogger(),
	}
}
