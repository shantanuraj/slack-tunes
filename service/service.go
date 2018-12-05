package service

import (
	"log"
	"time"

	"github.com/shantanuraj/slack-tunes/provider"
	"github.com/shantanuraj/slack-tunes/upstream"
)

// Service takes a `Provider` and `Upstream` and updates the `Upstream` at a fixed
// interval with the current song from `Provider`
type Service interface {
	Start(provider provider.Provider, upstream upstream.Upstream) error
}

// Sync implements `Service`
type Sync struct {
	updateInterval time.Duration
}

// Start starts the song-status sync service
func (s Sync) Start(p provider.Provider, u upstream.Upstream) error {
	var song provider.Song
	var err error
	for {
		if song, err = p.CurrentSong(); err != nil {
			log.Fatal("Couldn't get current song from provider", err)
		}

		if err = u.UpdateSong(song); err != nil {
			log.Fatal("Could not update upstream", err)
		}
		time.Sleep(s.updateInterval)
	}
}

// NewSync returns an instance of the sync service
func NewSync(updateInterval time.Duration) *Sync {
	return &Sync{
		updateInterval: updateInterval,
	}
}
