package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/blackironj/rest-be-template/model"
)

const _userCollectionName = "users"

type userRepository struct {
	collection *mongo.Collection
}

func newUsersRepository(db *mongo.Database) *userRepository {
	return &userRepository{
		collection: db.Collection(_userCollectionName),
	}
}

func (thiz *userRepository) UpsertUser(user *model.User, sc ...mongo.SessionContext) error {
	ctx := context.Background()
	if len(sc) > 0 {
		ctx = sc[0]
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"uid": user.UID}
	if _, err := thiz.collection.UpdateOne(ctx, filter, opts); err != nil {
		return err
	}
	return nil
}

func (thiz *userRepository) GetSingleUserByFilter(filter bson.M) (*model.User, error) {
	var user model.User
	if err := thiz.collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
