package response

import (
	. "goa.design/goa/v3/dsl"
)

var FindProductResult = ResultType("FindProductResult", func() {
	Field(1, "ID", UInt)
	Field(2, "UserID", UInt)
	Field(3, "Name", String)
	Field(4, "Author", String)
	Field(5, "Publisher", String)
	Field(6, "Price", Int)
	Field(7, "RegistrationDate", String)
})
