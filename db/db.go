package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// sqlite3 is the driver name and api.db is the data source name
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Println("Failed to connect to database:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF	NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		emailid TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		is_admin BOOLEAN DEFAULT FALSE,
		created_on TIMESTAMP,
		is_active TINYINT DEFAULT 1
	);`

	_, usersTableCreationErr := DB.Exec(createUsersTable)
	if usersTableCreationErr != nil {
		log.Println("Failed to create users table:", usersTableCreationErr)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_name TEXT NOT NULL,
		event_description TEXT NOT NULL,
		event_location TEXT NOT NULL,
		event_datetime DATETIME NOT NULL,
		user_id INTEGER,
		created_on TIMESTAMP,
		is_active TINYINT DEFAULT 1,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, eventsTableCreationErr := DB.Exec(createEventsTable)
	if eventsTableCreationErr != nil {
		log.Println("Failed to create events table:", eventsTableCreationErr)
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (event_id) REFERENCES events(id)
	);`

	_, registrationsTableCreationErr := DB.Exec(createRegistrationsTable)
	if registrationsTableCreationErr != nil {
		log.Println("Failed to create registrations table:", registrationsTableCreationErr)
	}
}
