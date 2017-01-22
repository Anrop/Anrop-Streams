package twitch

import (
	twitchAPI "github.com/knspriggs/go-twitch"
)

var (
	TwitchSession *twitchAPI.Session
)

func CreateSession(clientID string) error {
	var err error
	TwitchSession, err = twitchAPI.NewSession(twitchAPI.NewSessionInput{ClientID: clientID})
	return err
}
