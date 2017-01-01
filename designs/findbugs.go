package designs

import (
	_ "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// FindBugs represents all configurables for the FindBugs static checker
var FindBugs = Type("Findbugs", func() {
	Description("Description of which Findbugs bugs are checked.")

	Attribute("name")
	// TODO fill in with the list of FindBugs elements
})
