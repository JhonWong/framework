package main

import (
	"jwwebframework/framework/middleware"
)

func registerRoute(core *framework.Core) {
	core.Get("/user/login", middleware.Test3(), UserLoginControllerHandler)

	topGroup := core.Group("/top")
	topGroup.Use(middleware.Test3())

	subjectApi := topGroup.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelControllerHandler)
		subjectApi.Put("/:id", SubjectUpdateControllerHandler)
		subjectApi.Get("/:id", SubjectGetControllerHandler)
		subjectApi.Get("/list/all", SubjectListControllerHandler)
	}
}
