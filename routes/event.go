package routes

import (
	"log"
	"net/http"
	"strconv"

	event "example.com/event-booking-app/models"
	"github.com/gin-gonic/gin"
)

func saveEventData(context *gin.Context) {
	var event event.Event
	reqPayloadErr := context.ShouldBindJSON(&event)

	if reqPayloadErr != nil {
		log.Println("Error in request data payload =-=-=-=", reqPayloadErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid data received"})
		return
	} else {
		event.UserID = context.GetInt64("userId")
		eventSaveErr := event.Save()

		if eventSaveErr != nil {
			log.Println("Error in saving event =-=-=-=-=-=-=-=", eventSaveErr)
			context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to save data"})
			return
		}

		context.JSON(http.StatusCreated, gin.H{"status": true, "message": "Data saved successfully!"})
	}
}

func getAllEventsData(context *gin.Context) {
	var events, eventsErr = event.GetAllEventsData()

	if eventsErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", eventsErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in fetching events"})
		return
	}

	if len(events) <= 0 {
		context.JSON(http.StatusOK, gin.H{"status": false, "message": "No data found!", "data": events})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": true, "data": events})
	}
}

func getEventDataById(context *gin.Context) {
	idStr := context.Param("id")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var event, eventErr = event.GetEventDataById(idConv)

	if eventErr != nil {
		log.Println("Error in fetching data =-=-=-=-=-=-=-=", eventErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in fetching events"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": true, "data": event})
}

func updateEventDataById(context *gin.Context) {
	idStr := context.Param("id")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var event event.Event
	eventErr := context.ShouldBindJSON(&event)

	if eventErr != nil {
		log.Println("Error in event data payload =-=-=-=", eventErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid data received"})
		return
	}

	userId := context.GetInt64("userId")
	isAdmin := context.GetBool("isAdmin")

	if event.UserID != userId || isAdmin {
		context.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
		return
	}

	updateErr := event.UpdateEventDataById(idConv)

	if updateErr != nil {
		log.Println("Error in updating data =-=-=-=-=-=-=-=", updateErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in updating event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": true, "message": "Event updated successfully!"})
}

func deleteEventDataById(context *gin.Context) {
	idStr := context.Param("id")

	idConv, idConvErr := strconv.ParseInt(idStr, 10, 64)

	if idConvErr != nil {
		log.Println("Conversion of Id failed =-=-=-=-=-=-=", idConvErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid datatype for id"})
		return
	}

	var event event.Event

	var userIdReq map[string]int64
	reqErr := context.ShouldBindJSON(&userIdReq)

	if reqErr != nil {
		log.Println("Error in event data payload =-=-=-=", reqErr)
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid data received"})
		return
	}

	userId := context.GetInt64("userId")
	isAdmin := context.GetBool("isAdmin")

	if userIdReq["userId"] != userId || isAdmin {
		context.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
		return
	}

	deleteErr := event.DeleteEventDataById(idConv)

	if deleteErr != nil {
		log.Println("Error in deleting data =-=-=-=-=-=-=-=", deleteErr)
		context.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error in deleting event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": true, "message": "Event deleted successfully!"})
}
