package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "supersecret"

func GenerateToken(emailId string, userId int64, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"emailId": emailId,
		"userId":  userId,
		"isAdmin": isAdmin,
		"expiry":  time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, bool, error) {
	parsedToken, parsedTokenErr := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if parsedTokenErr != nil {
		log.Println("Error in parsing token =-=-=-=-=-=", parsedTokenErr)
		return 0, false, errors.New("could not parse token")
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, false, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, false, errors.New("claim token error")
	}

	// emailId := claims["emailId"].(string)
	userId := claims["userId"].(float64)
	isAdmin := claims["isAdmin"].(bool)

	return int64(userId), isAdmin, nil
}
