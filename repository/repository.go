package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/blackironj/rest-be-template/env"
)

var (
	mongoDB *mongo.Database

	userRepo *userRepository
)

func Init() {
	mongoDB = OpenMongoDB(env.MongoDBUrl, env.MongoDBname)

	userRepo = newUsersRepository(mongoDB)
}

func User() *userRepository {
	return userRepo
}
