package designs

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("gh", func() {
	BasePath("/gh")
	Description("For GH logins")

	Action("login", func() {
		Response(OK)
		Routing(POST("/"))
	})

	Action("login2", func() {
		Response(OK)
		Routing(GET("/"))
	})
})
