package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("bug_profile", func() {
	BasePath("/bug-profile")
	Description("Bug profile represents the list of bugs the user wants us to check for")

	// Corresponds to viewing a bug profile
	Action("show", func() {
		Description("Show a single bug profile")
		Routing(GET("/:profileID"))

		Params(func() {
			Param("profileID", Integer, "The profile's unique identifier")
		})

		Response(OK)
		Response(NotFound)
	})

	// Corresponds to listing all bug profiles
	Action("list", func() {
		Description("Show all of my bug profiles")
		Routing(GET("/"))

		Response(OK)
	})

	// Corresponds to removing a bug profile
	Action("delete", func() {
		Description("I don't want this bug profile anymore")
		Routing(DELETE("/:profileID"))

		Params(func() {
			Param("profileID", Integer, "The profile's unique identifier")
		})

		Response(OK)
		Response(NotFound)
	})

	// Corresponds to making a new bug profile
	Action("create", func() {
		Description("Make me a new bug profile")
		Routing(POST("/"))

		Payload(BugProfile)

		Response(OK)
	})

	// Corresponds to adjusting a pre-existing bug profile
	Action("update", func() {
		Description("Update my pre-existing bug profile")
		Routing(PUT("/:profileID"))

		Params(func() {
			Param("profileID", Integer, "The profile's unique identifier")
		})

		// TODO put in payload information
		Payload(BugProfile)

		Response(OK)
		Response(NotFound)
	})
})

// BugProfile represents the entire "FindBugs" and "Checkstyle" profiles
var BugProfile = Type("bug_profile", func() {
	Description("A bug profile")

	Attribute("name")
	Attribute("checkstyle", Checkstyle)
	Attribute("findbugs", FindBugs)
})
