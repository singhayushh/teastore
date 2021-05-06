package middlewares

import (
	"teastore/api/auth"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware checks whether user is signed in or not.
func AuthenticationMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		SessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(401, gin.H{"error": "Error accessing cookie"})
			c.Abort()
			return
		}

		uid, utype, err := auth.CheckSession(SessionID)
		if err != nil {
			c.JSON(401, gin.H{"error": "Dead Session"})
			c.Abort()
			return
		}
		if role == "Admin" && utype != role {
			c.JSON(401, gin.H{"error": "Not Admin"})
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}

// PasserMiddleware is the opposite of AuthenticationMiddleware
func PasserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SessionID, err := c.Cookie("session_id")
		if err != nil {
			c.Next()
		}

		_, _, err = auth.CheckSession(SessionID)
		if err != nil {
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "You are already logged in"})
			c.Abort()
		}
	}
}
