package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/blackironj/rest-be-template/model"
)

const _bookCollectionName = "books"

type bookRepository struct {
	collection *mongo.Collection
}

func newBooksRepository(db *mongo.Database) *bookRepository {
	return &bookRepository{
		collection: db.Collection(_bookCollectionName),
	}
}

func (thiz *bookRepository) UpsertBook(book *model.Book, sc ...mongo.SessionContext) error {
	ctx := context.Background()
	if len(sc) > 0 {
		ctx = sc[0]
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"isbn": book.ISBN}
	if _, err := thiz.collection.UpdateOne(ctx, filter, opts); err != nil {
		return err
	}
	return nil
}

func (thiz *bookRepository) GetSingleBookByFilter(filter bson.M) (*model.Book, error) {
	var book model.Book
	if err := thiz.collection.FindOne(context.Background(), filter).Decode(&book); err != nil {
		return nil, err
	}
	return &book, nil
}
