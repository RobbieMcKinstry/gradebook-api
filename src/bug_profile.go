package main

import (
	"github.com/alligrader/gradebook-api/src/app"
	"github.com/goadesign/goa"
)

// BugProfileController implements the bug_profile resource.
type BugProfileController struct {
	*goa.Controller
}

// NewBugProfileController creates a bug_profile controller.
func NewBugProfileController(service *goa.Service) *BugProfileController {
	return &BugProfileController{Controller: service.NewController("BugProfileController")}
}

// Create runs the create action.
func (c *BugProfileController) Create(ctx *app.CreateBugProfileContext) error {
	// BugProfileController_Create: start_implement

	// Put your logic here

	// BugProfileController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *BugProfileController) Delete(ctx *app.DeleteBugProfileContext) error {
	// BugProfileController_Delete: start_implement

	// Put your logic here

	// BugProfileController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *BugProfileController) List(ctx *app.ListBugProfileContext) error {
	// BugProfileController_List: start_implement

	// Put your logic here

	// BugProfileController_List: end_implement
	return nil
}

// Show runs the show action.
func (c *BugProfileController) Show(ctx *app.ShowBugProfileContext) error {
	// BugProfileController_Show: start_implement

	// Put your logic here

	// BugProfileController_Show: end_implement
	return nil
}

// Update runs the update action.
func (c *BugProfileController) Update(ctx *app.UpdateBugProfileContext) error {
	// BugProfileController_Update: start_implement

	// Put your logic here

	// BugProfileController_Update: end_implement
	return nil
}
