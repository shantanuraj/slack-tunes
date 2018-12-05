package upstream

import (
	"fmt"

	"github.com/shantanuraj/slack-tunes/provider"
)

// Slack updates the song info to a slack status
type Slack struct{}

// UpdateSong posts the song data to Slack's status API
func (i Slack) UpdateSong(song provider.Song) error {
	fmt.Println("Currently playing", song.Title, "by", song.Artist)
	return nil
}
