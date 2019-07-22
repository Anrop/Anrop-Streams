package twitch

import (
	twitchAPI "github.com/knspriggs/go-twitch"
)

var (
	// TwitchSession holds the session for TwitchAPI access
	TwitchSession *twitchAPI.Session
)

// CreateSession creates a new session to the Twitch API
func CreateSession(clientID string) error {
	var err error
	TwitchSession, err = twitchAPI.NewSession(twitchAPI.NewSessionInput{ClientID: clientID})
	return err
}
