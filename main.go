package main

import (
	"net/http"

	"example.com/event-booking-app/db"
	"example.com/event-booking-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/ping", pingResponse)
	routes.RegisterRouter(server)
	server.Run(":8080")
}

func pingResponse(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": true, "message": "PONG"})
}
