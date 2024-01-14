package routes

import (
	"log"
	"net/http"
	"strconv"

	event "example.com/event-booking-app/models"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	idStr := context.Param("eventId")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var event, eventErr = event.GetEventDataById(idConv)

	if eventErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", eventErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong"})
		return
	}

	userId := context.GetInt64("userId")

	registrationErr := event.Register(userId)

	if registrationErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", eventErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in registering for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": true, "message": "Registered for event successfully!"})
}

func deleteRegistration(context *gin.Context) {
	idStr := context.Param("eventId")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var event, eventErr = event.GetEventDataById(idConv)

	if eventErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", eventErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong"})
		return
	}

	userId := context.GetInt64("userId")

	deleteRegistrationErr := event.DeleteRegistration(userId)

	if deleteRegistrationErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", deleteRegistrationErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in deleting registerion for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": true, "message": "Deleted registration for event successfully!"})
}
