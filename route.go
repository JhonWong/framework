// Copyright 2023 jhonwong.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/jhonwong/framework/framework/gin"
	"github.com/jhonwong/framework/framework/middleware"
)

func registerRoute(core *gin.Engine) {
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	topGroup := core.Group("/top")
	topGroup.Use(middleware.Test3())

	subjectApi := topGroup.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)
	}
}
