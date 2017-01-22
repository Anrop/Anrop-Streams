package api

import (
	"db"
)

type Stream struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

type Streamer struct {
	Stream Stream  `json:"stream"`
	User   db.User `json:"user"`
}
