package db

import (
	"fmt"
	"os"
)

func GetUsers() (*[]User, error) {
	rows, err := Database.Query(`
		SELECT
			user_avatar,
			user_id,
			user_name,
			user_twitch
		FROM
			fusion7_users
		WHERE
			user_lastvisit > (UNIX_TIMESTAMP(NOW()) - 30 * 24 * 60 * 60)
			AND
			user_status = 0
			AND
			user_twitch IS NOT NULL
			AND
			CHAR_LENGTH(user_twitch) > 0
		ORDER BY
			user_name
	`)

	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Avatar, &user.ID, &user.Username, &user.TwitchID)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading user from database: %q", err)
		} else {
			user = fixAvatarURL(user)
			users = append(users, user)
		}
	}

	return &users, nil
}

func fixAvatarURL(user User) User {
	user.Avatar = os.Getenv("AVATAR_BASE_URL") + user.Avatar
	return user
}
