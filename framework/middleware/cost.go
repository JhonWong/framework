package middleware

import (
	"log"
	"time"

	"github.com/jhonwong/framework/framework/gin"
)

func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost %v", c.Request.RequestURI, cost.Seconds())
	}
}
