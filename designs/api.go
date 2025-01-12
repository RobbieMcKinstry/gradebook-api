package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("GradebookAPI", func() { // "My API" is the name of the API used in docs
	Title("Gradebook")                // Documentation title
	Description("The next big thing") // Longer documentation description
	Host("localhost:8081")            // Host used by Swagger and clients
	Scheme("http")                    // HTTP scheme used by Swagger and clients
	BasePath("/api")                  // Base path to all API endpoints
	Consumes("application/json")      // Media types supported by the API
	Produces("application/json")      // Media types generated by the API

	var JWTAuthModel = JWTSecurity("jwt", func() {
		Description("For routes that require an authorization token")
		Header("Authorization")
		TokenURL("https://example.com/token")
		Scope("settings:write", "Can update settings")
		Scope("settings:read", "Can read settings")
		Scope("profiles:write", "Can update profiles")
		Scope("profiles:read", "Can read profiles")
		Scope("projects:write", "Can update projects")
		Scope("projects:read", "Can read projects")
	})

	_ = JWTAuthModel
})
