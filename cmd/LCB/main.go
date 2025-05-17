package main

import (
	"github.com/wang900115/LCB/internal/adapter/gorm"
	"github.com/wang900115/LCB/internal/adapter/redispool"
	"github.com/wang900115/LCB/internal/application/usecase"
	"github.com/wang900115/LCB/pkg/config"
	"github.com/wang900115/LCB/pkg/logger"
)

func main() {
	conf := config.NewConfig()

	redispool := redispool.NewRedisPool(redispool.NewOption(conf))
	zaplogger := logger.NewZapLogger(logger.NewOption(conf))
	mysql := gorm.NewMysql(gorm.NewOption(conf))

	userRepo := repository.NewUserRepository(mysql)
	channelRepo := repository.NewChannelRepository(mysql)
	tokenRepo := repository.NewTokenRepository(redispool, conf.GetDuration("jwt.expiration"))

	userUsecase := usecase.NewUserUsecase(userRepo)
	channelUsecase := usecase.NewChannelUsecase(channelRepo)
	tokenUsecase := usecase.NewTokenUsecase(tokenRepo)

}
