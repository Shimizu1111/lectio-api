package response

import (
	. "goa.design/goa/v3/dsl"
)

var FindProductResult = ResultType("FindProductResult", func() {
	Field(1, "BookID", UInt)
	Field(2, "UserID", UInt)
	Field(3, "BookName", String)
	Field(4, "Author", String)
	Field(5, "Publisher", String)
	Field(6, "Price", Int)
	Field(7, "RegistrationDate", String)
})

var FindAllProductResult = ResultType("FindAllProductResult", func() {
	Field(1, "Books", Any)
})
