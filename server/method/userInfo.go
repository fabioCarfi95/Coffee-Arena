package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func UserInfo(c *gin.Context) {
	var user User

	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// Process the user information
	fmt.Println("Username:", user.Username)
	fmt.Println("Password:", user.Password)

	// Validate the username and password (you can replace this with your own validation logic)
	if user.Username != "admin" || user.Password != "password" {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	claims := JWTClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte("your-secret-key")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Store the token in the MySQL database
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	// Insert the token into the database
	_, err = db.Exec("INSERT INTO tokens (username, token, timestamp) VALUES (?, ?, ?)", user.Username, signedToken, time.Now().Unix())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to store token in database"})
		return
	}

	c.JSON(200, gin.H{"token": signedToken})
}
