package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//const secretKey string = "secretkey"

func getSecretKey() string {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	return os.Getenv("SECRET_KEY")
}

// Middleware JWT - verify and authenticate token
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil {
			//c.JSON(http.StatusInternalServerError, gin.H{"message": "Error when taking the token"})
			c.Abort()
			//c.Redirect(http.StatusInternalServerError, "/")
			return
			//loginRedirect(http.StatusInternalServerError, c)
		}
		if tokenString == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{"message": "Token not given"})
			c.Abort()
			return
			//c.Redirect(http.StatusUnauthorized, "/")
			//loginRedirect(http.StatusUnauthorized, c)
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if err, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("sign method not valid: %v", err)
			}
			return []byte(getSecretKey()), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token invalid"})
			c.SetCookie("token", "", -1, "/", "localhost", false, false)
			c.Abort()
			return
			//loginRedirect(http.StatusUnauthorized, c)
		}
		c.Next()
	}
}

// func loginRedirect(statusCode int, c *gin.Context) {
// 	c.HTML(statusCode, "login.html", nil)
// }
