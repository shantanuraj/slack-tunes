package provider

import (
	"log"
)

// Song represents a song playing in any provider
type Song struct {
	Title  string
	Artist string
}

type providerMaker = func() Provider

var providersMap = map[string]providerMaker{
	"itunes":  NewITunes,
	"spotify": NewSpotify,
	"tidal":   NewTidal,
}

// Provider is a generic type representing a song provider
type Provider interface {
	GetName() string
	// CurrentSong returns a currently playing song
	CurrentSong() (Song, error)
	// IsPlaying returns boolean to indicate if provider is playing
	IsPlaying() (bool, error)
}

// IsSameSong returns true if given songs are the same
func IsSameSong(a, b Song) bool {
	return a.Title == b.Title && a.Artist == b.Artist
}

// GetProvider returns the actual provivder for a given provider name
func GetProvider(providerName string) Provider {
	maker, ok := providersMap[providerName]

	if !ok {
		log.Fatal("Invalid provider specified")
	}

	return maker()
}
