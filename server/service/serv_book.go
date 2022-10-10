package service

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/blackironj/rest-be-template/model"
	repo "github.com/blackironj/rest-be-template/repository"
)

func RegisterBook(book *model.Book) error {
	return repo.Book().UpsertBook(book)
}

func GetBookByISBN(isbn string) *model.Book {
	filter := bson.M{
		"isbn": isbn,
	}
	book, err := repo.Book().GetSingleBookByFilter(filter)
	if err != nil {
		return nil
	}
	return book
}
