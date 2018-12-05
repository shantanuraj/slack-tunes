package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/shantanuraj/slack-tunes/logger"
	"github.com/shantanuraj/slack-tunes/provider"
	"github.com/shantanuraj/slack-tunes/service"
	"github.com/shantanuraj/slack-tunes/upstream"

	"github.com/urfave/cli"
)

const (
	// AppVersion is the verison of the app
	AppVersion = "1.2.0"
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
	// FlagVerbose is the flag name for turning on verbose mode
	FlagVerbose = "verbose"
	// FlagInterval is the flag name for specifying update interval
	FlagInterval = "interval"
	// DefaultInterval is the default interval value
	DefaultInterval = 10
)

// flags for app
var flags = []cli.Flag{
	cli.StringFlag{
		Name:  parseName(FlagProvider),
		Value: DefaultProvider,
		Usage: "provider to fetch song info from can be `itunes`, `spotify`",
	},
	cli.StringFlag{
		Name:  parseName(FlagUpstream),
		Value: DefaultUpstream,
		Usage: "upstream to update song info to can be `slack`",
	},
	cli.BoolFlag{
		Name:  FlagVerbose,
		Usage: "turn on logs",
	},
	cli.IntFlag{
		Name:  parseName(FlagInterval),
		Value: DefaultInterval,
		Usage: "set custom time interval between api updates in seconds",
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

func timeInSeconds(updateInterval int) time.Duration {
	return time.Duration(updateInterval) * time.Second
}

func run(c *cli.Context) error {
	l := logger.NewLogger(c.Bool(FlagVerbose))

	providerName := c.String(FlagProvider)
	upstreamName := c.String(FlagUpstream)
	updateInterval := c.Int(FlagInterval)

	l.Log(
		"[main]",
		"provider:",
		providerName,
		"upstream:",
		upstreamName,
		"update-interval:",
		strconv.Itoa(updateInterval)+"s",
	)

	p := provider.GetProvider(providerName)
	u := upstream.GetUpstream(upstreamName)
	s := service.NewSync(timeInSeconds(updateInterval))
	return s.Start(p, u)
}
