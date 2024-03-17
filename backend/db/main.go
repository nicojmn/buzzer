package db

import (
	"database/sql"
	"log"
	 _ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func InitDB() (int, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err);
		return -1, err;
	}
	
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS admin
	(id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	username PRIMARY KEY TEXT NOT NULL,
	password TEXT NOT NULL,
	admin TEXT);
	
	CREATE TABLE IF NOT EXISTS teams
	(id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	name TEXT NOT NULL,
	score INTEGER NOT NULL,
	`)

	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return 0, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(hash), nil
}

func AddUser(username string, password string, admin string) (int, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err);
		return -1, err;
	}

	// check if user exists
	// TODO complete check
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username)
	log.Println(row)




	password, err = hashPassword(password)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	_, err = db.Exec("INSERT INTO users (username, password, admin) VALUES (?, ?, ?)", username, password, admin)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return 0, nil
}