package entities

type TokenClaims struct {
	Username    string `json:"user_name"`
	Channelname string `json:"channel_name"`

	ExpiredAt int64 `json:"expired_at"`
}
