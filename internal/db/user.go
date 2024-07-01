package db

import (
	"log"

	"github.com/Sohail-9098/ms-assessment-app/internal/user"
)

func GetUserId(userId int) user.User {
	conn := connect()
	defer close(conn)

	query := "SELECT * FROM assessment_user WHERE user_id=$1;"

	rows, err := conn.Query(query, userId)
	if err != nil {
		log.Fatalf("error fetching question from db: %v", err)
	}

	var user user.User
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.UserName)
		if err != nil {
			log.Fatalf("error reading result set: %v", err)
		}
	}

	return user
}
