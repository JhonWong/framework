package middleware

import (
	"fmt"
	"jwwebframework/framework"
)

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()
		c.Next()

		return nil
	}
}

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
