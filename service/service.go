package service

import (
	"log"
	"time"

	"github.com/shantanuraj/slack-tunes/provider"
	"github.com/shantanuraj/slack-tunes/upstream"
)

var timeDelta = 10 * time.Second

// Service takes a `Provider` and `Upstream` and updates the `Upstream` at a fixed
// interval with the current song from `Provider`
type Service interface {
	Start(provider provider.Provider, upstream upstream.Upstream) error
}

// Sync implements `Service`
type Sync struct {
}

// Start starts the song-status sync service
func (a Sync) Start(p provider.Provider, u upstream.Upstream) error {
	var song provider.Song
	var err error
	for {
		if song, err = p.CurrentSong(); err != nil {
			log.Fatal("Couldn't get current song from provider", err)
		}

		if err = u.UpdateSong(song); err != nil {
			log.Fatal("Could not update upstream", err)
		}
		time.Sleep(timeDelta)
	}
	// return err
}
