# slack-tunes

Simple daemon to track and post currently playing song from iTunes to Slack.

Needs an environment variable `STUNES_SLACK_TOKEN`.
Get it from https://api.slack.com/custom-integrations/legacy-tokens

## Build

```shell
# Compile
make
```

## Install

```shell
# Install to GOPATH
make install
```

## Usage

```
go run cmd/main.go

# or
make install
slack-tunes
```
