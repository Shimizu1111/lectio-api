package repository

import (
	// "inventory-system/api/model/database"
	// "inventory-system/api/model/database/entity"
	// inventory "inventory-system/gen/inventory"

	"fmt"
	"log"

	"github.com/Shimizu1111/lectio/design/entity"
	lectio "github.com/Shimizu1111/lectio/design/gen/lectio"
	"github.com/Shimizu1111/lectio/src/database"
)

func SaveProduct(p *lectio.CreateProductPayload) {
	product := entity.Book{
		UserID:    p.UserID,
		Name:      p.BookName,
		Author:    *p.Author,
		Publisher: *p.Publisher,
		Price:     *p.Price,
	}
	database.Db.Create(&product)
}

func UpdateProduct(p *lectio.UpdateProductPayload) {
	var product entity.Book
	if err := database.Db.Where("id = ?", p).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	// update
	product.Name = *p.BookName
	product.Author = *p.Author
	product.Publisher = *p.Publisher
	product.Price = *p.Price
	database.Db.Save(&product)
}

func FindProduct(id uint) entity.Book {
	var product entity.Book
	fmt.Println("エラーの確認")
	if err := database.Db.Where("id = ?", id).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	return product
}

func FindAllProduct(products []*entity.Book) []*entity.Book {
	if err := database.Db.Find(&products).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}

	return products
}

func DeleteProduct(id int) {
	var product entity.Book
	if err := database.Db.Where("id = ?", id).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	database.Db.Delete(&product)
}
