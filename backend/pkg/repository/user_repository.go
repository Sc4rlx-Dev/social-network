package repository

import (
	"database/sql"
	"github.com/asoudri/social-network/pkg/models"
	_ "github.com/mattn/go-sqlite3"
)

func CreateUser(user models.User) error {
	db, err := sql.Open("sqlite3", "./social-network.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(email, nickname, first_name, last_name, password_hash, date_of_birth, gender) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, user.Nickname, user.FirstName, user.LastName, user.PasswordHash, user.DateOfBirth, user.Gender)
	return err
}

func GetUserByEmail(email string) (models.User, error) {
	db, err := sql.Open("sqlite3", "./social-network.db")
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow("SELECT id, email, password_hash FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	return user, err
}
