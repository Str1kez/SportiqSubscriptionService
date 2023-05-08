package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHeader struct {
	User string `header:"User" binding:"required,uuid"`
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := userHeader{}
		if err := c.ShouldBindHeader(&header); err != nil {
			details := []gin.H{{"msg": "Неверный заголовок пользователя", "type": "subscription.header.invalid"}}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"detail": details})
			return
		}
		c.Next()
	}
}
