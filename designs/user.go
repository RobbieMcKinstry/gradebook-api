package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/user")
	Description("Represents a user of the site")

	// Corresponds to creating a new user
	Action("create", func() {
		Description("Sign up for the first time")
		Routing(POST("/"))
		Payload(UserCreate)
		Response(OK, UserMedia)
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
