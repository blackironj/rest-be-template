package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/blackironj/rest-be-template/env"
)

var (
	mongoDB *mongo.Database

	bookRepo *bookRepository
)

func Init() {
	mongoDB = OpenMongoDB(env.MongoDBUrl, env.MongoDBname)

	bookRepo = newBooksRepository(mongoDB)
}

func Book() *bookRepository {
	return bookRepo
}
