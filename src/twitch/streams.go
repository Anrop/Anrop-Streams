package twitch

import (
	"db"
	"errors"
	"strings"

	twitchAPI "github.com/knspriggs/go-twitch"
)

func GetStreams(users []db.User) (*twitchAPI.GetStreamsOutputType, error) {
	var channels []string
	for _, user := range users {
		channels = append(channels, user.TwitchID)
	}

	if TwitchSession == nil {
		return nil, errors.New("Twitch Client not initialized")
	}

	streamsQuery := twitchAPI.GetStreamsInputType{Channel: strings.Join(channels, ",")}

	return TwitchSession.GetStream(&streamsQuery)
}
