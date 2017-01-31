package main

import (
	"github.com/alligrader/gradebook-api/src/app"
	"github.com/goadesign/goa"
)

// GithubTokenController implements the GithubToken resource.
type GithubTokenController struct {
	*goa.Controller
}

// NewGithubTokenController creates a GithubToken controller.
func NewGithubTokenController(service *goa.Service) *GithubTokenController {
	return &GithubTokenController{Controller: service.NewController("GithubTokenController")}
}

// Read runs the read action.
func (c *GithubTokenController) Read(ctx *app.ReadGithubTokenContext) error {
	// GithubTokenController_Read: start_implement

	// Put your logic here

	// GithubTokenController_Read: end_implement
	var token string = "9726f0609e3219e036a6490c3d72e13e2b3eb341"
	res := &app.GithubtokenMt{&token}
	return ctx.OK(res)
}
