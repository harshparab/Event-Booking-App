package routes

import (
	"example.com/event-booking-app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {

	// Events
	// Authenticated
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events/saveEventData", saveEventData)
	authenticated.PUT("/events/updateEventDataById/:id", updateEventDataById)
	authenticated.DELETE("/events/deleteEventDataById/:id", deleteEventDataById)

	// Registration for events
	authenticated.POST("/register/registration/:eventId", register)
	authenticated.DELETE("/register/deleteRegistration/:eventId", deleteRegistration)

	// Unauthenticated
	server.GET("/events/getAllEventsData", getAllEventsData)
	server.GET("/events/getEventDataById/:id", getEventDataById)

	// Users
	server.POST("/users/signUp", signUpUser)
	server.GET("/users/getAllUsersData", getAllUsersData)
	server.GET("/users/getUserDataById/:id", getUserDataById)
	server.PUT("/users/updateUserDataById/:id", updateUserDataById)
	server.DELETE("/users/deleteUserDataById/:id", deleteUserDataById)
	server.POST("/users/login", loginUser)
}
