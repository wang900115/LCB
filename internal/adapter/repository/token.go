package repository

import (
	"context"
	"crypto/rand"
	"errors"
	"time"

	"github.com/wang900115/LCB/internal/adapter/gorm/model"
	"github.com/wang900115/LCB/internal/domain/entities"
	"github.com/wang900115/LCB/internal/domain/irepository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

const (
	jwtsaltPrefix = "jwtsalt:"
	saltSize      = 16
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

type TokenRepository struct {
	redis      *redis.Client
	expiration time.Duration
}

func NewTokenRepository(redis *redis.Client, expiration time.Duration) irepository.TokenRepository {
	return &TokenRepository{
		redis:      redis,
		expiration: expiration,
	}
}

func (TokenRepository) GenerateSalt(saltSize int) []byte {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

func (r *TokenRepository) CreateToken(tokenClaims entities.TokenClaims) (string, error) {
	salt := r.GenerateSalt(saltSize)
	tokenClaimsModel := model.TokenClaims{
		Username:    tokenClaims.Username,
		Channelname: tokenClaims.Channelname,
	}
	tokenClaimsModel.ExpiresAt = jwt.NewNumericDate(time.Now().Add(r.expiration))

	_, err := r.redis.Set(context.Background(), jwtsaltPrefix+tokenClaims.Username+tokenClaims.Channelname, string(salt), r.expiration).Result()
	if err != nil {
		return "", err
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaimsModel).SignedString(append([]byte(tokenClaimsModel.Username+tokenClaims.Channelname), salt...))
}

func (r *TokenRepository) ValidateToken(token string) (entities.TokenClaims, error) {
	unvertifiedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return entities.TokenClaims{}, err
	}
	mapClaims, ok := unvertifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		return entities.TokenClaims{}, errors.New("token map failed")
	}

	userName, ok := mapClaims["user_name"].(string)
	if !ok {
		return entities.TokenClaims{}, errors.New("token map userName failed")
	}

	channelName, ok := mapClaims["channel_name"].(string)
	if !ok {
		return entities.TokenClaims{}, errors.New("token map channelName failed")
	}

	salt, err := r.redis.Get(context.Background(), jwtsaltPrefix+userName+channelName).Result()
	if err != nil {
		return entities.TokenClaims{}, err
	}

	key := []byte(userName + channelName + salt)
	tokenClaims, parseErr := jwt.ParseWithClaims(token, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if parseErr != nil {
		return entities.TokenClaims{}, parseErr
	}

	if !tokenClaims.Valid {
		return entities.TokenClaims{}, ErrTokenExpired
	}

	tokenClaimsModel, ok := tokenClaims.Claims.(*model.TokenClaims)
	if !ok {
		return entities.TokenClaims{}, ErrInvalidToken
	}
	return tokenClaimsModel.ToDomain(), nil
}

func (r *TokenRepository) DeleteToken(userName, channelName string) error {
	return r.redis.Del(context.Background(), jwtsaltPrefix+userName+channelName).Err()

}
