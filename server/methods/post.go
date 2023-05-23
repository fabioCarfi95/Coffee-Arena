package methods

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	fmt.Fprintln(c.Writer, "retrieve USER INFO")
	fmt.Println("Endpoint hit: userinfo")
}
