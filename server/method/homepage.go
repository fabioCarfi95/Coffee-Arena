package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint hit: homePage")
}
