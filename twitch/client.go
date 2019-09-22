package twitch

import (
	"github.com/nicklaw5/helix"
)

var (
	// Client holds the client for Twitch API access
	Client *helix.Client
)

// CreateSession creates a new session to the Twitch API
func CreateSession(clientID string) error {
	var err error
	Client, err = helix.NewClient(&helix.Options{
		ClientID: clientID,
	})
	return err
}
