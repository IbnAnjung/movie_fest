package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	"github.com/IbnAnjung/movie_fest/repository/redis/models"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/redis/go-redis/v9"
)

type userTokenRepository struct {
	client *redis.Client
}

func NewUserTokenRepository(client *redis.Client) enUser.UserTokenRepository {
	return &userTokenRepository{
		client: client,
	}
}

func (r *userTokenRepository) StoreToken(ctx *context.Context, userToken *enUser.UserToken) (err error) {
	m := models.UserToken{}
	m.FillFromEntity(*userToken)

	jm, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return r.client.Set(*ctx, m.Key(), jm, time.Until(userToken.Token.ExpiresAt)).Err()
}

func (r *userTokenRepository) GetToken(ctx *context.Context, userToken *enUser.UserToken) (err error) {
	m := models.UserToken{}
	m.FillFromEntity(*userToken)

	val, err := r.client.Get(*ctx, m.Key()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			e := utils.DataNotFoundError
			e.Message = "Token Not Found"
			err = e
		}

		return err
	}

	fmt.Println(val)
	if err = json.Unmarshal([]byte(val), &m); err != nil {
		return err
	}

	m.ToEntity(userToken)
	return
}

func (r *userTokenRepository) DeleteToken(ctx *context.Context, id string) (err error) {
	m := models.UserToken{}
	m.TokenID = id
	return r.client.Del(*ctx, m.Key()).Err()
}
