package entities

type User struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`

	ChannelID uint `json:"channel_id"`
}
