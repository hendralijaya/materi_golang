package service

import (
	"materi-golang/exception"
	"materi-golang/model/domain"
	"materi-golang/model/web"
	"materi-golang/repository"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b web.BookCreateRequest) (domain.Book,error)
	Update(b web.BookUpdateRequest) (domain.Book, error)
	Delete(bookId uint64) error
	FindById(bookId uint64) (domain.Book, error)
	All() []domain.Book
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService (bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository: bookRepository}
}

func (service *bookService) Insert(request web.BookCreateRequest) (domain.Book, error) {
	book := domain.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&request))
	if(err != nil) {
		return book, err
	}
	return service.bookRepository.Insert(book), nil
}

func (service *bookService) Update(request web.BookUpdateRequest) (domain.Book, error) {
	bookRequest := domain.Book{}
	err := smapping.FillStruct(&bookRequest, smapping.MapFields(&request))
	if(err != nil) {
		return bookRequest, err
	}
	_ ,err = service.bookRepository.FindById(request.Id)
	if err != nil {
		return bookRequest, exception.NewNotFoundError(err.Error())
	}
	return service.bookRepository.Update(bookRequest), nil
}

func (service *bookService) Delete(bookId uint64) error {
	book , err := service.bookRepository.FindById(bookId)
	if err != nil {
		return exception.NewNotFoundError(err.Error())
	}
	service.bookRepository.Delete(book)
	return nil
}

func (service *bookService) FindById(bookId uint64) (domain.Book, error) {
	book, err := service.bookRepository.FindById(bookId)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (service *bookService) All() []domain.Book {
	return service.bookRepository.All()
}