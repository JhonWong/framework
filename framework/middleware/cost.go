// Copyright 2023 jhonwong.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
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
