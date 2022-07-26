package repository

import (
	// "inventory-system/api/model/database"
	// "inventory-system/api/model/database/entity"
	// inventory "inventory-system/gen/inventory"

	"fmt"
	"log"
	"time"

	"github.com/Shimizu1111/lectio/design/entity"
	lectio "github.com/Shimizu1111/lectio/design/gen/lectio"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	*gorm.DB
}

func (s *Repository) SaveProduct(p *lectio.CreateProductPayload) {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	product := entity.Book{
		UserID:           p.UserID,
		Name:             p.BookName,
		Author:           *p.Author,
		Publisher:        *p.Publisher,
		Price:            *p.Price,
		RegistrationDate: nowTime,
		BookStatus:       *p.BookStatus,
		CreatedAt:        nowTime,
		UpdatedAt:        nowTime,
	}
	fmt.Println(product)
	s.Create(&product)
}

func (s *Repository) UpdateProduct(p *lectio.UpdateProductPayload) {
	var product entity.Book
	if err := s.Where("id = ?", p.BookID).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	// update
	product.Name = *p.BookName
	product.Author = *p.Author
	product.Publisher = *p.Publisher
	product.Price = *p.Price
	s.Save(&product)
}

func (s *Repository) FindProduct(id uint) entity.Book {
	var product entity.Book
	fmt.Println("エラーの確認")
	if err := s.Where("id = ?", id).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	return product
}

func (s *Repository) FindAllProduct(products []*entity.Book) []*entity.Book {
	if err := s.Find(&products).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}

	return products
}

func (s *Repository) DeleteProduct(id int) {
	var product entity.Book
	if err := s.Where("id = ?", id).First(&product).Error; err != nil {
		log.Fatal("NOT FOUND PRODUCT")
	}
	s.Delete(&product)
}
