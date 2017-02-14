package main

import (
	"log"

	"github.com/alligrader/gradebook-api/src/app"
	"github.com/goadesign/goa"
)

// GhController implements the gh resource.
type GhController struct {
	*goa.Controller
}

// NewGhController creates a gh controller.
func NewGhController(service *goa.Service) *GhController {
	return &GhController{Controller: service.NewController("GhController")}
}

// Login runs the login action.
func (c *GhController) Login(ctx *app.LoginGhContext) error {
	// GhController_Login: start_implement

	// Put your logic here

	// GhController_Login: end_implement
	return nil
}

// Login2 runs the login2 action.
func (c *GhController) Login2(ctx *app.Login2GhContext) error {
	// GhController_Login2: start_implement

	// Put your logic here
	log.Println("We did it boys!")

	// GhController_Login2: end_implement
	return nil
}
