package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/rafimuhammad01/auth-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	mongoDB *mongo.Database
	redis   *redis.Client
}

func NewUser(mongoDB *mongo.Database, redis *redis.Client) *User {
	return &User{
		mongoDB: mongoDB,
		redis:   redis,
	}
}

func (u *User) GetUserDataFromToken(refreshToken string) (map[string]interface{}, []error) {
	ctx := context.Background()

	val, err := u.redis.Get(ctx, refreshToken).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, []error{fmt.Errorf("[%w] refresh token %s is not found or expired", domain.ErrRefreshTokenInvalid, refreshToken)}
		}
		return nil, []error{fmt.Errorf("[%w] internal server error from redis %s", domain.ErrInternal, err.Error())}
	}

	var res map[string]interface{}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return nil, []error{fmt.Errorf("[%w] internal server error from redis %s", domain.ErrInternal, err.Error())}
	}

	return res, nil
}

func (u *User) SaveToken(userid string, role int, refreshToken string) []error {
	ctx := context.Background()

	bs, _ := json.Marshal(map[string]interface{}{
		"user_id": userid,
		"role":    role,
	})

	status := u.redis.Set(ctx, refreshToken, bs, 168*time.Hour)

	if status.Err() != nil {
		return []error{fmt.Errorf("[%w] internal server error from redis %s", domain.ErrInternal, status.Err().Error())}
	}

	return nil
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
