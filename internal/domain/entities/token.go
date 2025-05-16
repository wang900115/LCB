package entities

type TokenClaims struct {
	UserID    uint `json:"user_id"`
	ChannelID uint `json:"channel_id"`

	Username    string `json:"user_name"`
	Channelname string `json:"channel_name"`

	ExpiredAt int64 `json:"expired_at"`
}
