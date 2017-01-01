package designs

import (
	_ "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Checkstyle contains a full specification of which Checkstyle lints are to be checked.
var Checkstyle = Type("Checkstyle", func() {
	Description("Description of which Checkstyle lints are checked.")

	// TODO fill in with the list of Checkstyle attributes
	// Remove the Name attribute!
	Attribute("name")
})
