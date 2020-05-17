package twitch

import (
	"errors"
	"github.com/nicklaw5/helix"
)

// GetStreams returns a list of channels from twitch
func GetStreams(channels []string) (*[]helix.Stream, error) {
	if Client == nil {
		return nil, errors.New("TwitchClient not initialized")
	}

	err := UpdateAppToken()
	if err != nil {
		return nil, err
	}

	response, err := Client.GetStreams(&helix.StreamsParams{
		UserLogins: channels,
	})

	if err != nil {
		return nil, err
	}

	return &response.Data.Streams, nil
}
