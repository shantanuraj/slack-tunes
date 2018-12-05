package provider

// Tidal app name
const Tidal = "TIDAL"

// NewTidal returns an instance of the `Tidal` provider
func NewTidal() Provider {
	// Doesn't work atm
	return NewMacProvider(Tidal)
}
