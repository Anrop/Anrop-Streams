# Anrop Streams API

Returns active streamers on Anrop.

## Requirements

Code is written in [Go](https://golang.org/) and uses
[Go modules](https://github.com/golang/go/wiki/Modules) for dependency
management.

## How To Use

Download all dependencies and compile the sources with `go build`.

Start the API with the `Anrop-Streams` binary.
Server will be available at `$PORT`.

## Environment Variables

Environment variables can be specified in `.env` file and will be autoloaded.

| Key | Required | Description |
| --- | -------- | ----------- |
| AVATAR_BASE_URL | No | Prefix to all avatar urls |
| DATABASE_URL | Yes | MySQL url to your Anrop database |
| TWITCH_CLIENT_ID | Yes | Twitch API Client ID |
| PORT | No | Port that HTTP Server is bound to. 8080 by default |
