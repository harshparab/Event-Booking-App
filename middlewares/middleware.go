package middlewares

import (
	"log"
	"net/http"

	"example.com/event-booking-app/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		log.Println("Token is empty")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
		return
	}

	userId, isAdmin, tokenVerificationErr := utils.VerifyToken(token)

	if tokenVerificationErr != nil {
		log.Println("Token verification failed =-=-=-=-=-=", tokenVerificationErr)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
		return
	}

	context.Set("userId", userId)
	context.Set("isAdmin", isAdmin)
	context.Next()
}
