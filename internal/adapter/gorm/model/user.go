package model

import (
	"github.com/wang900115/LCB/internal/domain/entities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string  `json:"name" gorm:"not null;unique"`
	Channel Channel `gorm:"foreignKey:ChannelID; references:ID"`
}

func (u User) TableName() string {
	return "users"
}

func (u *User) ToDomain() entities.User {
	return entities.User{
		ID:        u.ID,
		Name:      u.Name,
		ChannelID: u.Channel.ID,
	}
}
