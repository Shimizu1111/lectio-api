package design

import (
	"context"
	"log"

	"github.com/Shimizu1111/lectio/design/entity"
	lectio "github.com/Shimizu1111/lectio/design/gen/lectio"
	"github.com/Shimizu1111/lectio/src/domain/repository"
	"github.com/jinzhu/copier"
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
	s.logger.Print("lectio.create")
	return
}

// Update implements update.
func (s *lectiosrvc) Update(ctx context.Context, p *lectio.UpdateProductPayload) (res string, err error) {
	s.logger.Print("lectio.update")
	return
}

// Find implements find.
func (s *lectiosrvc) Find(ctx context.Context, p *lectio.FindProductPayload) (res *lectio.Findproductresult, err error) {
	res = &lectio.Findproductresult{}
	s.logger.Print("lectio.find")
	return
}

// FindAll implements findAll.
func (s *lectiosrvc) FindAll(ctx context.Context) (res []*lectio.Findproductresult, err error) {
	var books []*entity.Book
	products := repository.FindAllProduct(books)
	var result_books []*lectio.Findproductresult
	copier.Copy(&result_books, &products)
	s.logger.Print("lectio.findAll")
	return result_books, nil
}

// Delete implements delete.
func (s *lectiosrvc) Delete(ctx context.Context, p *lectio.DeletePayload) (res string, err error) {
	s.logger.Print("lectio.delete")
	return
}
