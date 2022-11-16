package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/rafimuhammad01/user-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	mongoDB *mongo.Database
}

func NewUser(mongoDB *mongo.Database) *User {
	return &User{
		mongoDB: mongoDB,
	}
}

func (u *User) GetAll() ([]*domain.User, []error) {
	ctx := context.Background()

	cur, err := u.mongoDB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	var result []*domain.User
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	return result, nil
}

func (u *User) GetByID(id string) (*domain.User, []error) {
	ctx := context.Background()
	col := u.mongoDB.Collection("users").FindOne(ctx, bson.M{"_id": id})

	var result domain.User
	err := col.Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, []error{fmt.Errorf("[%w] user with id %s not found", domain.ErrUserNotFound, id)}
		}
	}

	return &result, nil
}

func (u *User) GetByUsername(username string) (*domain.User, []error) {
	ctx := context.Background()
	col := u.mongoDB.Collection("users").FindOne(ctx, bson.M{"username": username})

	var result domain.User
	err := col.Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, []error{fmt.Errorf("[%w] user with username %s not found", domain.ErrUserNotFound, username)}
		}
	}

	return &result, nil
}

func (u *User) Create(user domain.User) []error {
	ctx := context.Background()

	_, err := u.mongoDB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	return nil
}

func (u *User) Update(user domain.User) []error {
	ctx := context.Background()
	filter := bson.M{"_id": user.ID}

	pByte, err := bson.Marshal(user)
	if err != nil {
		return []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	_, err = u.mongoDB.Collection("users").UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	return nil
}

func (u *User) Delete(id string) []error {
	ctx := context.Background()
	filter := bson.M{"_id": id}

	_, err := u.mongoDB.Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		return []error{fmt.Errorf("[%w] internal mongo error %s", domain.ErrUserInternal, err.Error())}
	}

	return nil
}
