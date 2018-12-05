package provider

const iTunes = "iTunes"

// NewITunes returns an instance of the `ITunes` provider
func NewITunes() Provider {
	return NewMacProvider(iTunes)
}
