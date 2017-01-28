package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("GithubToken", func() {

	BasePath("/GHtoken")
	Description("This is your GitHub API token")

	Action("read", func() {
		Description("Returns the GH token")
		Routing(GET("/"))

		Response(OK, GithubToken)
	})
})

var _ = Resource("user", func() {
	BasePath("/user")
	Description("Represents a user of the site")

	// Corresponds to creating a new user
	Action("create", func() {
		Description("Sign up for the first time")
		Routing(POST("/"))
		Payload(UserCreate)
		Response(OK, func() {
			Status(200)
			Media(UserMedia, "github")
		})
		Response(InternalServerError)
	})

	Action("update", func() {
		Description("Adjust your account settings")
		Routing(PUT("/:userID"))
		Payload(UserCreate)
		Response(OK, UserMedia)
		Response(InternalServerError)

		Params(func() {
			Param("userID", Integer, "The user's unique identifier")
		})
	})

	Action("read", func() {
		Description("Get the detials about this particular account")
		Routing(GET("/:userID"))
		Response(OK, UserMedia)
		Response(InternalServerError)

		Params(func() {
			Param("userID", Integer, "The user's unique identifier")
		})
	})
})

var GithubToken = MediaType("application/githubToken.mt", func() {
	Attributes(func() {
		Attribute("token")
	})

	View("default", func() {
		Attribute("token")
	})
})

var UserMedia = MediaType("application/user.mt", func() {
	Attributes(func() {
		Attribute("id", Integer)
		Attribute("name")
		Attribute("email", String, "User email", func() {
			Format("email")
		})
		Attribute("phone_number")
		Attribute("callback")
	})
	View("default", func() {
		Attribute("id", Integer)
		Attribute("name")
		Attribute("email", String, "User email", func() {
			Format("email")
		})
		Attribute("phone_number")
	})

	View("github", func() {
		Attribute("id", Integer)
		Attribute("name")
		Attribute("email", String, "User email", func() {
			Format("email")
		})
		Attribute("phone_number")
		Attribute("callback")
	})
})

var UserCreate = Type("user", func() {
	Description("payload used to sign up a user")
	Attribute("name")
	Attribute("email", String, "User email", func() {
		Format("email")
	})
	Attribute("phone_number")
	Attribute("password")
})
