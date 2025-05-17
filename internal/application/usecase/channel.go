package usecase

import (
	"github.com/wang900115/LCB/internal/domain/entities"
	"github.com/wang900115/LCB/internal/domain/irepository"
)

type ChannelUsecase struct {
	ChannelRepo irepository.ChannelRepository
}

func NewChannelUsecase(channelRepo irepository.ChannelRepository) *ChannelUsecase {
	return &ChannelUsecase{ChannelRepo: channelRepo}
}

func (c *ChannelUsecase) CreateChannel(channel entities.Channel) (entities.Channel, error) {
	return c.ChannelRepo.CreateChannel(channel)
}

func (c *ChannelUsecase) QueryChannel() ([]entities.Channel, error) {
	return c.ChannelRepo.QueryChannel()
}

func (c *ChannelUsecase) QueryUsers(channelName string) ([]entities.User, error) {
	return c.ChannelRepo.QueryUsers(channelName)
}
