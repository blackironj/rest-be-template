package service

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/blackironj/rest-be-template/model"
	repo "github.com/blackironj/rest-be-template/repository"
)

func GetUserByEmail(email string) *model.User {
	filter := bson.M{
		"email": email,
	}
	user, err := repo.User().GetSingleUserByFilter(filter)
	if err != nil {
		return nil
	}
	return user
}
