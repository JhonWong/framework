package main

import (
	"jwwebframework/framework"
	"time"
)

func UserLoginControllerHandler(c *framework.Context) error {
	time.Sleep(10 * time.Second)
	c.SetStatus(200).Json("ok, UserLoginControllerHandler")
	return nil
}

func SubjectDelControllerHandler(c *framework.Context) error {
	c.SetStatus(200).Json("ok SubjectDelControllerHandler")
	return nil
}

func SubjectUpdateControllerHandler(c *framework.Context) error {
	c.SetStatus(200).Json("ok SubjectUpdateControllerHandler")
	return nil
}

func SubjectGetControllerHandler(c *framework.Context) error {
	c.SetStatus(200).Json("ok SubjectGetControllerHandler")
	return nil
}

func SubjectListControllerHandler(c *framework.Context) error {
	c.SetStatus(200).Json("ok SubjectListControllerHandler")
	return nil
}
