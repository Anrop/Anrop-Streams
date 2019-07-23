package db

// User database model
type User struct {
	Avatar   string `json:"avatar"`
	ID       string `json:"id"`
	Username string `json:"username"`
	TwitchID string `json:"twitch_id"`
}
