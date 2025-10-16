package repository

import (
	"database/sql"
	// "time"
	"github.com/asoudri/social-network/pkg/models"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSession(session models.Session) error {
	db, err := sql.Open("sqlite3", "./social-network.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO sessions(user_id, token, expires_at) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(session.UserID, session.Token, session.ExpiresAt)
	return err
}