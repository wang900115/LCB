package irepository

import "github.com/wang900115/LCB/internal/domain/entities"

type ChannelRepository interface {
	CreateChannel(entities.Channel) (entities.Channel, error)
	QueryChannel() ([]entities.Channel, error)
	QueryUsers(string) ([]entities.User, error)
}
