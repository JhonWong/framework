// Copyright 2023 jhonwong.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"time"

	"github.com/jhonwong/framework/framework/gin"
)

func UserLoginController(c *gin.Context) {
	time.Sleep(10 * time.Second)
	c.ISetStatus(200).IJson("ok, UserLoginControllerHandler")
}

func SubjectDelController(c *gin.Context) {
	c.ISetStatus(200).IJson("ok SubjectDelControllerHandler")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetStatus(200).IJson("ok SubjectUpdateControllerHandler")
}

func SubjectGetController(c *gin.Context) {
	c.ISetStatus(200).IJson("ok SubjectGetControllerHandler")
}

func SubjectListController(c *gin.Context) {
	c.ISetStatus(200).IJson("ok SubjectListControllerHandler")
}
