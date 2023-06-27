package middleware

import "github.com/jhonwong/framework/framework/gin"

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.ISetStatus(500).IJson(err)
			}
		}()
		c.Next()
	}
}
