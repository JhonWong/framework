package middleware

import (
	"context"
	"fmt"
	"jwwebframework/framework"
	"log"
	"time"
)

func Timeout() framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			time.Sleep(10 * time.Second)
			c.SetOkStatus().Json("ok")
			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			c.WriteMux().Lock()
			defer c.WriteMux().Unlock()
			log.Println(p)
			c.SetStatus(500).Json("panic")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.WriteMux().Lock()
			defer c.WriteMux().Unlock()
			c.SetStatus(500).Json("time out")
			c.SetTimeout()
		}

		return nil

	}
}
