package main

import (
	"github.com/alligrader/gradebook-api/src/app"
	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here

	// UserController_Create: end_implement
	res := &app.UserMt{}
	return ctx.OK(res)
}
