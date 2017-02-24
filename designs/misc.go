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

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("api/swagger.json", "src/swagger/swagger.json")
})

var _ = Resource("github_auth", func() {
	BasePath("/auth/gh")
	Description("Auth with GitHub OAuth")

	Action("create", func() {
		Description("Create or login a user with GitHub OAuth")
		// Need to accept a payload object ? What gets returned?
		// Need to respond with a redirect.
		Response(Found)
		Routing(POST("/"))
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
