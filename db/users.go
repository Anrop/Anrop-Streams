package db

import (
	"fmt"
	"os"
)

// GetUsers returns a list of all users in database that have a twitch setting
// and have been on the webpage for the last 30 days or returns error.
func GetUsers() (*[]User, error) {
	rows, err := Database.Query(`
		SELECT
			users.user_avatar,
			users.user_id,
			users.user_name,
			users.user_twitch
		FROM
			fusion7_users AS users
		LEFT JOIN fusion7_operations_slots AS slots ON slots.user_id = users.user_id
		WHERE
			users.user_lastvisit > (UNIX_TIMESTAMP(NOW()) - 30 * 24 * 60 * 60)
			AND
			users.user_status = 0
			AND
			users.user_twitch IS NOT NULL
			AND
			CHAR_LENGTH(users.user_twitch) > 0
		GROUP BY
			slots.user_id
		HAVING
			COUNT(slots.user_id) > 10
		ORDER BY
			users.user_name
	`)

	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Avatar, &user.ID, &user.Username, &user.TwitchID)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading user from database: %q\n", err)
		} else {
			user = fixAvatarURL(user)
			users = append(users, user)
		}
	}

	rows.Close()

	return &users, nil
}

func fixAvatarURL(user User) User {
	user.Avatar = os.Getenv("AVATAR_BASE_URL") + user.Avatar
	return user
}
