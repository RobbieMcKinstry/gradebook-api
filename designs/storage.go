package designs

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("GlobalStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {

		Model("User", func() {
			BuildsFrom(func() {
				Payload("user", "create")
			})
			RendersTo(UserMedia)
			Description("User Model Description")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
			})
		})
	})
})
