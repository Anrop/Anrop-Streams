package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Anrop/Anrop-Streams/db"
	"github.com/Anrop/Anrop-Streams/twitch"
)

// StreamsHandler used to handle API request for current streams
func StreamsHandler(w http.ResponseWriter, r *http.Request) {
	streamers, err := getStreamers()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(streamers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling json: %q\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func getStreamers() (*[]Streamer, error) {
	users, err := db.GetUsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading users from database: %q\n", err)
		return nil, err
	}

	streams, err := getTwitchStreams(*users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching data from Twitch: %q\n", err)
		return nil, err
	}

	streamers := make([]Streamer, 0, len(*streams))
	for _, user := range *users {
		stream, ok := (*streams)[strings.ToLower(user.TwitchID)]
		if ok {
			streamers = append(streamers, Streamer{Stream: stream, User: user})
		}
	}

	return &streamers, nil
}

func getTwitchStreams(users []db.User) (*map[string]Stream, error) {
	var channels []string
	for _, user := range users {
		channels = append(channels, user.TwitchID)
	}

	streams, err := twitch.GetStreams(channels)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching streams from Twitch: %q\n", err)
		return nil, err
	}

	streamsMap := make(map[string]Stream)
	for _, stream := range *streams {
		streamsMap[strings.ToLower(stream.UserName)] = Stream{
			Image: formatThumbnailURL(stream.ThumbnailURL),
			Link:  fmt.Sprintf("https://twitch.tv/%s", stream.UserName),
			Title: stream.Title,
		}
	}

	return &streamsMap, nil
}

func formatThumbnailURL(thumbnailURL string) string {
	// Old Large was 640x360
	return strings.Replace(strings.Replace(thumbnailURL, "{width}", "640", 1), "{height}", "360", 1)
}
