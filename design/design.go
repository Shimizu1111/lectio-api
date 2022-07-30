package design

import (
	"github.com/Shimizu1111/lectio/design/request"
	"github.com/Shimizu1111/lectio/design/response"
	. "goa.design/goa/v3/dsl"
)

// API 定義
var _ = API("lectio", func() {
	// API の説明（タイトルと説明）
	Title("Lectio-API")
	Description("About Books")

	// サーバ定義
	Server("lectio", func() {
		Host("localhost", func() {
			URI("http://localhost:8000") // HTTP REST API
			URI("grpc://localhost:8080") // gRPC
		})
	})
})

// サービス定義
var _ = Service("lectio", func() {
	Description("About Book")

	// Create
	Method("create", func() {
		Payload(request.CreateProductPayload)
		Result(String)
		HTTP(func() {
			POST("/v1/books")
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
	// Update
	Method("update", func() {
		Payload(request.UpdateProductPayload)
		Result(String)
		HTTP(func() {
			PUT("/v1/books/{bookId}")
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
	// Find
	Method("find", func() {
		Payload(request.FindProductPayload)
		Result(response.FindProductResult)
		HTTP(func() {
			GET("/v1/books/{bookId}")
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
	// FindAll
	Method("findAll", func() {
		Result(ArrayOf(response.FindProductResult))
		HTTP(func() {
			GET("/v1/books")
		})
	})
	// Delete
	Method("delete", func() {
		Payload(func() {
			Field(1, "bookId", Int)
			Required("bookId")
		})
		Result(String)
		HTTP(func() {
			DELETE("/v1/books/{bookId}")
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})
