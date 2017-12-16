package twitch

import (
	"errors"
	"strings"

	twitchAPI "github.com/knspriggs/go-twitch"
)

func GetStreams(channels []string) (*twitchAPI.GetStreamsOutputType, error) {
	if TwitchSession == nil {
		return nil, errors.New("Twitch Client not initialized")
	}

	streamsQuery := twitchAPI.GetStreamsInputType{Channel: strings.Join(channels, ",")}

	return TwitchSession.GetStream(&streamsQuery)
}
