package design

import (
	"context"
	"fmt"
	"log"

	lectio "github.com/Shimizu1111/lectio/design/gen/lectio"
	"github.com/Shimizu1111/lectio/src/domain/repository"
)

// lectio service example implementation.
// The example methods log the requests and return zero values.
type lectiosrvc struct {
	logger *log.Logger
}

// NewLectio returns the lectio service implementation.
func NewLectio(logger *log.Logger) lectio.Service {
	return &lectiosrvc{logger}
}

// Create implements create.
func (s *lectiosrvc) Create(ctx context.Context, p *lectio.CreateProductPayload) (res string, err error) {
	repository.SaveProduct(p)
	s.logger.Print("lectio.create")
	return
}

// Update implements update.
func (s *lectiosrvc) Update(ctx context.Context, p *lectio.UpdateProductPayload) (res string, err error) {
	repository.UpdateProduct(p)
	s.logger.Print("lectio.update")
	return
}

// Find implements find.
func (s *lectiosrvc) Find(ctx context.Context, p *lectio.FindProductPayload) (res *lectio.Findproductresult, err error) {
	product := repository.FindProduct(p.BookID)
	res = &lectio.Findproductresult{
		BookID:           &p.BookID,
		UserID:           &product.UserID,
		BookName:         &product.BookName,
		Author:           &product.Author,
		Publisher:        &product.Publisher,
		Price:            &product.Price,
		RegistrationDate: &product.RegistrationDate,
	}
	s.logger.Print("lectio.find")
	return
}

// FindAll implements findAll.
func (s *lectiosrvc) FindAll(ctx context.Context) (res *lectio.Findallproductresult, err error) {
	products := repository.FindAllProduct()
	// var product entity.Book
	// if err := database.Db.Find(&product); err != nil {
	// 	fmt.Printf("p :%+v\n", err)
	// 	//log.Fatal(&err)
	// }
	// var result *lectio.Findallproductresult
	// res = products
	//var result entity.Books
	//res.Books = []*entity.Book{}
	//entity.Books = []*entity.Book{}

	//for i := 0; i < len(products); i++ {
	// res = &lectio.Findallproductresult{
	// 	Books: entity.Books{
	// 		entity.Book{
	// 			BookID:           products.BookID,
	// 			UserID:           &products[i].UserID,
	// 			BookName:         &products[i].BookName,
	// 			Author:           &products[i].Author,
	// 			Publisher:        &products[i].Publisher,
	// 			Price:            &products[i].Price,
	// 			RegistrationDate: &products[i].RegistrationDate,
	// 		},
	// 	},
	// }
	//}
	var product = repository.FindProduct(p.ProductID)
	res = &inventory.Findproductresult{
		ProductID:          &p.ProductID,
		ProductName:        &product.ProductName,
		ProductDescription: &product.ProductDescription,
		ProductMinStock:    &product.ProductMinStock,
	}

	// res = append(result)
	// fmt.Println("要素数の取得3")
	// fmt.Println(*res)
	//}

	//aaa = &lectio.Findallproductresult
	fmt.Println("要素数の取得")
	fmt.Println(products)
	// res = &lectio.Findallproductresult{
	// 	BookID:           &products[1].BookID,
	// 	UserID:           &products[1].UserID,
	// 	BookName:         &products[1].BookName,
	// 	Author:           &products[1].Author,
	// 	Publisher:        &products[1].Publisher,
	// 	Price:            &products[1].Price,
	// 	RegistrationDate: &products[1].RegistrationDate,
	// }
	s.logger.Print("lectio.findAll")
	return
}

// Delete implements delete.
func (s *lectiosrvc) Delete(ctx context.Context, p *lectio.DeletePayload) (res string, err error) {
	repository.DeleteProduct(p.BookID)
	s.logger.Print("lectio.delete")
	return
}
