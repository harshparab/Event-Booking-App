package models

import (
	"log"

	"example.com/event-booking-app/db"
)

func (e *Event) Register(userId int64) error {
	saveQuery := `
		INSERT INTO registrations (event_id, user_id) VALUES (?, ?);
	`
	prepQuery, prepErr := db.DB.Prepare(saveQuery)

	if prepErr != nil {
		log.Println("Prepare statement error =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	response, responseErr := prepQuery.Exec(e.ID, userId)

	if responseErr != nil {
		log.Println("Query execution error =-=-=-=-=-=", responseErr)
		return responseErr
	}

	lastId, lastIdErr := response.LastInsertId()

	if lastIdErr != nil {
		log.Println("Failed to insert data =-=-=-=-=-=", lastIdErr)
		return lastIdErr
	}

	e.ID = lastId

	return nil
}

func (e *Event) DeleteRegistration(userId int64) error {
	deleteSingleQuery := `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?;
	`

	prepQuery, prepErr := db.DB.Prepare(deleteSingleQuery)

	if prepErr != nil {
		log.Println("Error in query =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	_, responseErr := prepQuery.Exec(e.ID, userId)

	if responseErr != nil {
		log.Println("Failed to delete data =-=-=-=-=-=-=", responseErr)
		return responseErr
	}

	return nil
}
