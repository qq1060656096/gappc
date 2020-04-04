package generate

const MiddlewareTemplate = `package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func {{middlewareName}}() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		logrus.Info(">>>>>>....", token)
		// // validate fail
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		c.Set("token", token)
		c.Next()
	}
}
`