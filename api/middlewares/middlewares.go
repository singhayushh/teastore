package middlewares

import (
	"fmt"
	"teastore/api/auth"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware checks whether user is signed in or not.
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(401, gin.H{"error": "Error accessing cookie"})
			fmt.Println(err)
			c.Abort()
			return
		}

		Email, err := auth.CheckSession(SessionID)
		if err != nil {
			c.JSON(401, gin.H{"error": "Dead Session"})
			c.Abort()
			return
		}
		fmt.Println(Email)
		c.Next()
	}
}
