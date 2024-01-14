package models

import (
	"log"
	"time"

	"example.com/event-booking-app/db"
)

type Event struct {
	ID               int64
	EventName        string    `binding:"required"`
	EventDescription string    `binding:"required"`
	EventLocation    string    `binding:"required"`
	EventDateTime    time.Time `binding:"required"`
	UserID           int64
	CreatedOn        time.Time
}

func (e *Event) Save() error {
	saveQuery := `
		INSERT INTO events (event_name, event_description, event_location, event_datetime, user_id) VALUES (?, ?, ?, ?, ?);
	`
	prepQuery, prepErr := db.DB.Prepare(saveQuery)

	if prepErr != nil {
		log.Println("Prepare statement error =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	response, responseErr := prepQuery.Exec(e.EventName, e.EventDescription, e.EventLocation, e.EventDateTime, e.UserID)

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

func GetAllEventsData() ([]Event, error) {
	fetchAllEventsQuery := `
		SELECT id, event_name, event_description, event_location, event_datetime, user_id FROM events;
	`
	response, responseErr := db.DB.Query(fetchAllEventsQuery)

	if responseErr != nil {
		log.Println("Failed to fetch data =-=-=-=-=-=", responseErr)
		return nil, responseErr
	}

	var events []Event

	for response.Next() {
		var event Event
		response.Scan(&event.ID, &event.EventName, &event.EventDescription, &event.EventLocation, &event.EventDateTime, &event.UserID)
		events = append(events, event)
	}

	return events, nil
}

func GetEventDataById(id int64) (*Event, error) {
	fetchSingleEventQuery := `
		SELECT id, event_name, event_description, event_location, event_datetime, user_id FROM events where id = ?;
	`
	response := db.DB.QueryRow(fetchSingleEventQuery, id)

	var event Event

	scanErr := response.Scan(&event.ID, &event.EventName, &event.EventDescription, &event.EventLocation, &event.EventDateTime, &event.UserID)

	if scanErr != nil {
		log.Println("Error in scanning id =-=-=-=-=-=-=", scanErr)
		return nil, scanErr
	}

	return &event, nil
}

func (e *Event) UpdateEventDataById(id int64) error {
	updateSingleQuery := `
		UPDATE events SET event_name = ?, event_description = ?, event_location = ?, event_datetime = ? WHERE id = ?;
	`

	prepQuery, prepErr := db.DB.Prepare(updateSingleQuery)

	if prepErr != nil {
		log.Println("Error in query =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	_, responseErr := prepQuery.Exec(e.EventName, e.EventDescription, e.EventLocation, e.EventDateTime, id)

	if responseErr != nil {
		log.Println("Failed to update data =-=-=-=-=-=-=", responseErr)
		return responseErr
	}

	return nil
}

func (e *Event) DeleteEventDataById(id int64) error {
	deleteSingleQuery := `
		DELETE FROM events WHERE id = ?;
	`

	prepQuery, prepErr := db.DB.Prepare(deleteSingleQuery)

	if prepErr != nil {
		log.Println("Error in query =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	_, responseErr := prepQuery.Exec(id)

	if responseErr != nil {
		log.Println("Failed to delete data =-=-=-=-=-=-=", responseErr)
		return responseErr
	}

	return nil
}
