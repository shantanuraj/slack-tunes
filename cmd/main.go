package main

import (
	"fmt"
	"os"

	"github.com/shantanuraj/slack-tunes/provider"
	"github.com/shantanuraj/slack-tunes/service"
	"github.com/shantanuraj/slack-tunes/upstream"

	"github.com/urfave/cli"
)

const (
	// AppVersion is the verison of the app
	AppVersion = "0.0.1"
	// AppName is the name of the app
	AppName = "slack-tunes"
	// AppUsage is the help text for the app
	AppUsage = "Post currently playing song to your status"
	// FlagProvider defines the flag name for specifying the provider
	FlagProvider = "provider"
	// DefaultProvider is the default provider
	DefaultProvider = "itunes"
	// FlagUpstream defines the flag name for specifying the upstream
	FlagUpstream = "upstream"
	// DefaultUpstream is the default upstream
	DefaultUpstream = "slack"
)

// flags for app
var flags = []cli.Flag{
	cli.StringFlag{
		Name:  parseName(FlagProvider),
		Value: DefaultProvider,
		Usage: "provider to fetch song info from can be `itunes`",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion
	app.Flags = flags
	app.Action = run
	app.Run(os.Args)
}

func parseName(flagName string) string {
	return fmt.Sprintf("%s, %v", flagName, string(flagName[0]))
}

func run(c *cli.Context) error {
	p := provider.GetProvider(c.String(FlagProvider))
	u := upstream.GetUpstream(c.String(FlagUpstream))
	s := service.Sync{}
	return s.Start(p, u)
}
