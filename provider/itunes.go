package provider

// ITunes fetches song info from iTunes
type ITunes struct{}

// CurrentSong fetches the current song from iTunes
func (i ITunes) CurrentSong() (Song, error) {
	song := Song{
		Title:  "I Am",
		Artist: "Jorja Smith",
	}
	return song, nil
}

// NewITunes returns an instance of the `ITunes` provider
func NewITunes() ITunes {
	return ITunes{}
}
