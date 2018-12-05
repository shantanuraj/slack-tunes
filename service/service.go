package service

import (
	"github.com/shantanuraj/slack-tunes/provider"
	"github.com/shantanuraj/slack-tunes/upstream"
)

// Service takes a `Provider` and `Upstream` and updates the `Upstream` at a fixed
// interval with the current song from `Provider`
type Service interface {
	Start(provider provider.Provider, upstream upstream.Upstream) error
}
