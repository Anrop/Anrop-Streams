package api

import (
	"github.com/Anrop/Anrop-Streams/db"
)

// Stream data model
type Stream struct {
	Image string `json:"image"`
	Link  string `json:"link"`
	Title string `json:"title"`
}

// Streamer data model
type Streamer struct {
	Stream Stream  `json:"stream"`
	User   db.User `json:"user"`
}
