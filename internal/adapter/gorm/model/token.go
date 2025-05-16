package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/wang900115/LCB/internal/domain/entities"
)

type TokenClaims struct {
	UserID    uint `json:"user_id"`
	ChannelID uint `json:"channel_id"`

	Username    string `json:"user_name"`
	Channelname string `json:"channel_name"`

	jwt.RegisteredClaims
}

func (t TokenClaims) ToDomain() entities.TokenClaims {
	return entities.TokenClaims{
		UserID:    t.UserID,
		ChannelID: t.ChannelID,

		Username:    t.Username,
		Channelname: t.Channelname,

		ExpiredAt: t.ExpiresAt.Unix(),
	}
}
