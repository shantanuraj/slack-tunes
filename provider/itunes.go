package provider

const iTunes = "Music"

// NewITunes returns an instance of the `ITunes` provider
func NewITunes() Provider {
	return NewMacProvider(iTunes)
}
