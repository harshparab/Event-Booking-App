package routes

import (
	"log"
	"net/http"
	"strconv"

	user "example.com/event-booking-app/models"
	"example.com/event-booking-app/utils"
	"github.com/gin-gonic/gin"
)

func signUpUser(context *gin.Context) {
	var user user.User
	reqPayloadErr := context.ShouldBindJSON(&user)

	if reqPayloadErr != nil {
		log.Println("Error in request payload =-=-=-=", reqPayloadErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid data received"})
		return
	} else {
		userSaveErr := user.Save()

		if userSaveErr != nil {
			log.Println("Error in saving event =-=-=-=-=-=-=-=", userSaveErr)
			context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to save data"})
			return
		}

		context.JSON(http.StatusCreated, gin.H{"status": true, "message": "Data saved successfully!"})
	}
}

func getAllUsersData(context *gin.Context) {
	var users, usersErr = user.GetAllUsersData()

	if usersErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", usersErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in fetching users"})
		return
	}

	if len(users) <= 0 {
		context.JSON(http.StatusOK, gin.H{"status": false, "message": "No data found!", "data": users})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": true, "data": users})
	}
}

func getUserDataById(context *gin.Context) {
	idStr := context.Param("id")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var user, userErr = user.GetUserDataById(idConv)

	if userErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", userErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in fetching user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": true, "data": user})
}

func updateUserDataById(context *gin.Context) {
	idStr := context.Param("id")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var user user.User
	reqPayloadErr := context.ShouldBindJSON(&user)

	if reqPayloadErr != nil {
		log.Println("Error in request data payload =-=-=-=", reqPayloadErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid data received"})
		return
	}

	updateErr := user.UpdateUserDataById(idConv)

	if updateErr != nil {
		log.Println("Error in updating data =-=-=-=-=-=-=-=", updateErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in updating user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": true, "message": "User updated successfully!"})
}

func deleteUserDataById(context *gin.Context) {
	idStr := context.Param("id")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	deleteErr := user.DeleteUserDataById(idConv)

	if deleteErr != nil {
		log.Println("Error in deleting data =-=-=-=-=-=-=-=", deleteErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in deleting user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": true, "message": "User deleted successfully!"})
}

func loginUser(context *gin.Context) {
	var user user.User
	reqPayloadErr := context.ShouldBind(&user)

	if reqPayloadErr != nil {
		log.Println("Error in request payload =-=-=-=", reqPayloadErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid data received"})
		return
	} else {
		validateCredentials := user.LoginUser()

		if validateCredentials != nil {
			log.Println("Invalid login credentials =-=-=-=-=-=", validateCredentials)
			context.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": error.Error(validateCredentials)})
			return
		}

		token, tokenErr := utils.GenerateToken(user.EmailId, user.ID, user.IsAdmin)

		if tokenErr != nil {
			log.Println("Error in generating token =-=-=-=-=-=-=-=", tokenErr)
			context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to authenticate user"})
			return
		}

		context.JSON(http.StatusAccepted, gin.H{"status": true, "message": "Login successful!", "userId": user.ID, "isAdmin": user.IsAdmin, "token": token})
	}
}
