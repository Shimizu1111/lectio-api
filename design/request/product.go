package request

import (
	. "goa.design/goa/v3/dsl"
)

var CreateProductPayload = Type("CreateProductPayload", func() {
	Field(1, "userId", UInt)
	Field(2, "bookName", String)
	Field(3, "author", String)
	Field(4, "publisher", String)
	Field(5, "price", Int)
	Field(6, "registrationDate", String)
	Field(7, "bookStatus", String)
	Required("userId", "bookName")
})

var UpdateProductPayload = Type("UpdateProductPayload", func() {
	Field(1, "bookId", UInt)
	Field(2, "bookName", String)
	Field(3, "author", String)
	Field(4, "publisher", String)
	Field(5, "price", Int)
	Field(6, "registrationDate", String)
	Required("bookId")
})

var FindProductPayload = Type("FindProductPayload", func() {
	Field(1, "bookId", UInt)
	Field(2, "userId", UInt)
	Field(3, "bookName", String)
	Field(4, "author", String)
	Field(5, "publisher", String)
	Field(6, "price", Int)
	Field(7, "registrationDate", String)
	Required("bookId")
})
