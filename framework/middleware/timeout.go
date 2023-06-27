// Copyright 2023 jhonwong.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jhonwong/framework/framework/gin"
)

func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1*time.Second))
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			time.Sleep(10 * time.Second)
			c.ISetOkStatus().IJson("ok")
			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			c.ISetStatus(500).IJson("panic")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.ISetStatus(500).IJson("time out")
		}
	}
}
