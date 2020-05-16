package twitch

import (
	"github.com/nicklaw5/helix"
	"time"
)

var (
	// Client holds the client for Twitch API access
	Client          *helix.Client
	token           *helix.AppAccessCredentials
	tokenExpiryTime time.Time
)

// CreateSession creates a new session to the Twitch API
func CreateSession(clientID string, clientSecret string) error {
	var err error
	Client, err = helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	return err
}

// UpdateAppToken updates client authorization with a new app token when needed
func UpdateAppToken() error {
	if token == nil || time.Now().After(tokenExpiryTime) {
		tokenRequestTime := time.Now()
		res, err := Client.GetAppAccessToken()
		if err != nil {
			return err
		}
		token = &res.Data
		tokenExpiryTime = tokenRequestTime.Add(time.Duration(token.ExpiresIn) * time.Second)
		Client.SetAppAccessToken(token.AccessToken)
	}
	return nil
}
