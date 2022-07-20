package repository

import (
	"errors"
	"materi-golang/model/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	All() []domain.Book
	Insert(b domain.Book) domain.Book
	Update(b domain.Book) domain.Book
	Delete(b domain.Book)
	FindById(bookId uint64) (domain.Book, error)
}

type BookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(connection *gorm.DB) BookRepository {
	return &BookConnection{connection: connection}
}

func(db *BookConnection) All() []domain.Book {
	var books []domain.Book
	db.connection.Find(&books)
	return books
}

func (db *BookConnection) Insert(book domain.Book) domain.Book {
	db.connection.Save(&book)
	db.connection.Find(&book)
	return book
}

func (db *BookConnection) Update(book domain.Book) domain.Book {
	db.connection.Save(&book)
	db.connection.Find(&book)
	return book
}

func (db *BookConnection) Delete(book domain.Book) {
	db.connection.Delete(&book)
}

func (db *BookConnection) FindById(id uint64) (domain.Book, error) {
	var book domain.Book
	db.connection.Find(&book, id)
	if book.Id == 0 {
		return book, errors.New("book not found")
	}
	return book, nil
}