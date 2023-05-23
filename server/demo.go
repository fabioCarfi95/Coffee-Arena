package main

import (
	"fmt"
	"time"

	//"github.com/fabioCarfi95/Coffee-Arena/server/method/login"
	"github.com/gin-gonic/gin"
)

type NextSession struct {
	Time time.Time `json:"time"`
}

func nextSession(c *gin.Context) {
	fmt.Println("Endpoint hit: nextSession")
	c.JSON(200, gin.H{"next-session": time.Now().Unix()})
}



