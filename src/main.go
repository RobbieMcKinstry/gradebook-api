//go:generate goagen bootstrap -d github.com/alligrader/gradebook-api/designs

package main

import (
	"github.com/alligrader/gradebook-api/src/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("GradebookAPI")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "bug_profile" controller
	c := NewBugProfileController(service)
	app.MountBugProfileController(service, c)
	// Mount "user" controller
	c2 := NewUserController(service)
	app.MountUserController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}
}
