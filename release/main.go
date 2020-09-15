package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/urfave/cli"
)

const (
	// EnvArch defines enviornment var name for OS
	EnvArch = "GOARCH"
	// EnvOS defines enviornment var name for OS
	EnvOS = "GOOS"
	// EnvPath defines environment path for Go
	EnvPath = "GOPATH"

	// FlagArch defines flag name for arch option
	FlagArch = "arch"
	// FlagBaseDir defines flag name for base directory option
	FlagBaseDir = "base-dir"
	// FlagOS defines flag name for os option
	FlagOS = "os"
)

// flags for build tool
var flags = []cli.Flag{
	cli.StringFlag{
		Name:  FlagOS,
		Value: runtime.GOOS,
		Usage: "Operating system to build for",
	},
	cli.StringFlag{
		Name:  FlagArch,
		Value: runtime.GOARCH,
		Usage: "Architecture to build for",
	},
	cli.StringFlag{
		Name:  FlagBaseDir,
		Value: "",
		Usage: "Installation directory",
	},
}

func main() {
	app := configureApp()
	app.Run(os.Args)
}

// configureApp configures the build cli app
func configureApp() *cli.App {
	app := cli.NewApp()

	app.Name = "slack-tunes"
	app.Usage = "Build slack-tunes tool"
	app.Flags = flags
	app.Action = build

	return app
}

// build compiles the slack-tunes tool for given configuration
func build(c *cli.Context) error {
	buildOS := c.String(FlagOS)
	buildArch := c.String(FlagArch)
	baseDir := c.String(FlagBaseDir)
	goPath := os.Getenv(EnvPath)

	outName := "slack-tunes"
	if baseDir == "" {
		outName = formatName(outName, buildOS, buildArch)
	}

	inFile := path.Join("cmd", "slack-tunes.go")
	outFile := path.Join(baseDir, "bin", outName)

	args := []string{"build", "-o", outFile, inFile}

	cmd := exec.Command("go", args...)

	cmd.Env = setEnv(os.Environ(), EnvOS, buildOS)
	cmd.Env = setEnv(cmd.Env, EnvArch, buildArch)
	cmd.Env = setEnv(cmd.Env, EnvPath, goPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running", strings.Join(cmd.Args, " "))
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error building binary: %v\n", err)
	} else {
		fmt.Printf("Built and installed to %s\n", outFile)
	}

	return err
}

// Create an environment variable of the form key=value.
func envPair(key, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}

// setEnv sets the given key & value in the provided environment.
// Each value in the env list should be of the form key=value.
func setEnv(env []string, key, value string) []string {
	for i, s := range env {
		if strings.HasPrefix(s, fmt.Sprintf("%s=", key)) {
			env[i] = envPair(key, value)
			return env
		}
	}
	env = append(env, envPair(key, value))
	return env
}

// formatName returns the formatted name for release file
func formatName(name, os, arch string) string {
	return fmt.Sprintf("%s-%s-%s-%v", name, os, arch, time.Now().Unix())
}
