package models

import (
	"errors"
	"log"
	"time"

	"example.com/event-booking-app/db"
	"example.com/event-booking-app/utils"
)

type User struct {
	ID        int64
	EmailId   string `binding:"required"`
	Password  string `binding:"required"`
	IsAdmin   bool
	IsActive  int
	CreatedOn time.Time
}

func (u *User) Save() error {
	saveQuery := `
		INSERT INTO users (emailid, password) VALUES (?, ?);
	`
	prepQuery, prepErr := db.DB.Prepare(saveQuery)

	if prepErr != nil {
		log.Println("Prepare statement error =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	encryptedPassword, encryptionErr := utils.EncryptPassword(u.Password)

	if encryptionErr != nil {
		log.Println("Failed to encrypt password =-=-=-=-=-=-=", encryptionErr)
		return encryptionErr
	}

	response, responseErr := prepQuery.Exec(u.EmailId, encryptedPassword)

	if responseErr != nil {
		log.Println("Query execution error =-=-=-=-=-=", responseErr)
		return responseErr
	}

	lastId, lastIdErr := response.LastInsertId()

	if lastIdErr != nil {
		log.Println("Failed to insert data =-=-=-=-=-=", lastIdErr)
		return lastIdErr
	}

	u.ID = lastId

	return nil
}

func GetAllUsersData() ([]User, error) {
	fetchAllUsersQuery := `
		SELECT id, emailid, password, is_admin, is_active FROM users;
	`
	response, responseErr := db.DB.Query(fetchAllUsersQuery)

	if responseErr != nil {
		log.Println("Failed to fetch data =-=-=-=-=-=", responseErr)
		return nil, responseErr
	}

	var users []User

	for response.Next() {
		var user User
		response.Scan(&user.ID, &user.EmailId, &user.Password, &user.IsAdmin, &user.IsActive)
		users = append(users, user)
	}

	return users, nil
}

func GetUserDataById(id int64) (*User, error) {
	fetchSingleEventQuery := `
		SELECT id, emailid, password, is_admin, is_active FROM users where id = ?;
	`
	response := db.DB.QueryRow(fetchSingleEventQuery, id)

	var user User

	scanErr := response.Scan(&user.ID, &user.EmailId, &user.Password, &user.IsAdmin, &user.IsActive)

	if scanErr != nil {
		log.Println("Error in scanning id =-=-=-=-=-=-=", scanErr)
		return nil, scanErr
	}

	return &user, nil
}

func DeleteUserDataById(id int64) error {
	deleteSingleQuery := `
		DELETE FROM users WHERE id = ?;
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

func (u *User) UpdateUserDataById(id int64) error {
	updateSingleQuery := `
		UPDATE users SET emailid = ?, password = ?, is_admin = ?, is_active = ? WHERE id = ?;
	`

	prepQuery, prepErr := db.DB.Prepare(updateSingleQuery)

	if prepErr != nil {
		log.Println("Error in query =-=-=-=-=-=", prepErr)
		return prepErr
	}

	defer prepQuery.Close()

	encryptedPassword, encryptionErr := utils.EncryptPassword(u.Password)

	if encryptionErr != nil {
		log.Println("Failed to encrypt password =-=-=-=-=-=-=", encryptionErr)
		return encryptionErr
	}

	_, responseErr := prepQuery.Exec(u.EmailId, encryptedPassword, u.IsAdmin, u.IsActive, id)

	if responseErr != nil {
		log.Println("Failed to update data =-=-=-=-=-=-=", responseErr)
		return responseErr
	}

	return nil
}

func (u *User) LoginUser() error {
	fetchUserDataQuery := `
		SELECT id, password, is_admin, is_active FROM users WHERE is_active = 1 and emailid = ?;  
	`
	response := db.DB.QueryRow(fetchUserDataQuery, u.EmailId)

	var payloadPassword string

	responseScanErr := response.Scan(&u.ID, &payloadPassword, &u.IsAdmin, &u.IsActive)

	if responseScanErr != nil {
		log.Println("Error in scanning response =-=-=-=-=-=-=", responseScanErr)
		return errors.New("unregistered user")
	}

	isPasswordValid := utils.ComparePasswords(payloadPassword, u.Password)

	if !isPasswordValid {
		log.Println("Password does not match")
		return errors.New("invalid password")
	}

	return nil
}
