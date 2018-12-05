package provider

import (
	"errors"
	"os/exec"
	"strings"
)

const osascript = "osascript"
const tellCmd = "tell application"
const playerStateCmd = "player state"

// run executes the AppleScript command and returns the output
func run(command string) (string, error) {
	cmd := exec.Command(osascript, "-e", command)
	out, err := cmd.CombinedOutput()
	prettyOutput := strings.Replace(string(out), "\n", "", -1)

	if err != nil {
		return "", errors.New(err.Error() + ": " + prettyOutput + " (" + command + ")")
	}

	return prettyOutput, nil
}

// build the AS command, removing blanks
func build(args ...string) string {
	var params []string

	for _, arg := range args {
		if arg != "" {
			params = append(params, arg)
		}
	}

	return strings.Join(params, " ")
}

// buildTell builds the tell command
func buildTell(application string, commands ...string) string {
	app := "\"" + application + "\""
	args := []string{tellCmd, app, "\n"}
	for _, command := range commands {
		args = append(args, command, "\n")
	}
	args = append(args, "end", "tell")
	return build(args...)
}

// tell calls executes the tell subcommand for osascript
func tell(application string, commands ...string) (string, error) {
	return run(buildTell(application, commands...))
}
