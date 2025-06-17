package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/Aashish-star/Go_AcharyaPrashantAssignment/model"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("abcxyz")

func GenerateJWT(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 60).Unix(), // Token expires in 15 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Error while generating token ", err)
	}
	return "Bearer " + tokenStr
}

func GenerateRefreshToken(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Error while generating token ", err)
	}
	return "Bearer " + tokenStr
}

func GetUsernameFromToken(tokenIn string) string {

	tokenString := strings.TrimPrefix(tokenIn, "Bearer ")

	// Parse the token
	claims := &model.JwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	fmt.Println("Token err ", err)
	if err == nil && token.Valid {
		return claims.Username
	}
	return ""
}

func ValidateToken(token string) string {
	fmt.Println("Incoming token", token)
	if strings.HasPrefix(token, "Bearer ") {

		tokenString := strings.TrimPrefix(token, "Bearer ")
		// Parse the token

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Make sure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {

			if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
				return "Token expired"
			}
			fmt.Println("Error parsing token:", err)
			return "Invalid token"
		}

		if token.Valid {
			return "Token is valid"
		}
	}
	return "Header is invalid"

}
