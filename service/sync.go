package service

import (
	"log"
	"time"

	"github.com/shantanuraj/slack-tunes/logger"

	"github.com/shantanuraj/slack-tunes/provider"
	"github.com/shantanuraj/slack-tunes/upstream"
)

// Sync implements `Service`
type Sync struct {
	updateInterval time.Duration
	lastSong       provider.Song
	logger         *logger.Logger
}

// Reset the status
func (s *Sync) Reset(p provider.Provider, u upstream.Upstream) error {
	return u.UpdateSong(p.GetName(), false, provider.Song{})
}

// Start starts the song-status sync service
func (s *Sync) Start(p provider.Provider, u upstream.Upstream) error {
	var isPlaying bool
	var song provider.Song
	var err error
	providerName := p.GetName()
	for {
		isPlaying, err = p.IsPlaying()

		if err != nil {
			log.Fatal("Couldn't get playing state from provider", err)
		}

		if !isPlaying {
			s.logger.Log("[service-sync] Nothing playing currently")
			goto Update
		}

		if song, err = p.CurrentSong(); err != nil {
			log.Fatal("Couldn't get current song from provider", err)
		}
		s.logger.Log("[service-sync] Currently playing", song.Title, "by", song.Artist)

		s.lastSong = song
	Update:
		if err = u.UpdateSong(providerName, isPlaying, song); err != nil {
			log.Fatal("Could not update upstream", err)
		}

		time.Sleep(s.updateInterval)
	}
}

// NewSync returns an instance of the sync service
func NewSync(updateInterval time.Duration) *Sync {
	return &Sync{
		updateInterval: updateInterval,
		lastSong:       provider.Song{},
		logger:         logger.GetLogger(),
	}
}
