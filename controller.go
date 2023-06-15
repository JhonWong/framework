package main

import (
	"context"
	"fmt"
	"jwwebframework/framework"
	"time"
)

func UserLoginControllerHandler(c *framework.Context) error {
	c.Json(200, "ok, UserLoginControllerHandler")
	return nil
}

func SubjectDelControllerHandler(c *framework.Context) error {
	c.Json(200, "ok, SubjectDelControllerHandler")
	return nil
}

func SubjectUpdateControllerHandler(c *framework.Context) error {
	c.Json(200, "ok, SubjectUpdateControllerHandler")
	return nil
}

func SubjectGetControllerHandler(c *framework.Context) error {
	c.Json(200, "ok, SubjectGetControllerHandler")
	return nil
}

func SubjectListControllerHandler(c *framework.Context) error {
	c.Json(200, "ok, SubjectListControllerHandler")
	return nil
}

func FooControllerHandler(c *framework.Context) error {
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case <-panicChan:
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()
		//...
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()
		//...
		c.Json(500, "time out")
		c.SetTimeout()
	}

	return c.Json(200, map[string]interface{}{
		"code": 0,
	})
}
