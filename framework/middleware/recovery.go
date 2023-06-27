// Copyright 2023 jhonwong.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
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
