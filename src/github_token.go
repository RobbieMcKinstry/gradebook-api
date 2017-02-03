package main

import (
	"io/ioutil"
	"log"

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
	bytes, err := ioutil.ReadFile("./.deploy/github_api.txt")
	if err != nil {
		log.Println(err)
	}
	var token string = string(bytes)

	// GithubTokenController_Read: end_implement
	res := &app.GithubtokenMt{&token}
	return ctx.OK(res)
}
