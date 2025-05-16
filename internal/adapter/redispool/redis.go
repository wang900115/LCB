package redispool

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Option struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewOption(conf *viper.Viper) Option {
	return Option{
		Addr:     conf.GetString("redis.host"),
		Username: conf.GetString("redis.user"),
		Password: conf.GetString("redis.password"),
		DB:       conf.GetInt("redis.database"),
	}
}

func NewRedisPool(option Option) *redis.Client {
	redisPool := redis.NewClient(&redis.Options{
		Addr:     option.Addr,
		Username: option.Username,
		Password: option.Password,
		DB:       option.DB,
	})
	return redisPool
}
