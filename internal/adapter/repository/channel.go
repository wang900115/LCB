package repository

import (
	"github.com/wang900115/LCB/internal/adapter/gorm/model"
	"github.com/wang900115/LCB/internal/domain/entities"
	"gorm.io/gorm"
)

type ChannelRepository struct {
	gorm *gorm.DB
}

func (c *ChannelRepository) CreateChannel(channel entities.Channel) (entities.Channel, error) {
	channelModel := model.Channel{
		Name: channel.Name,
	}

	if err := c.gorm.Create(&channel).Error; err != nil {
		return entities.Channel{}, err
	}

	return channelModel.ToDomain(), nil
}

func (c *ChannelRepository) QueryChannels() ([]entities.Channel, error) {
	var channelsModel []model.Channel
	if err := c.gorm.Find(&channelsModel).Error; err != nil {
		return nil, err
	}

	var channels []entities.Channel
	for _, channelModel := range channelsModel {
		channels = append(channels, channelModel.ToDomain())
	}
	return channels, nil
}

func (c *ChannelRepository) QueryUsers(channelName string) ([]entities.User, error) {
	var channelModel model.Channel
	if err := c.gorm.Preload("Users").Where("name = ?", channelName).First(&channelModel).Error; err != nil {
		return nil, err
	}

	var users []entities.User
	for _, userModel := range channelModel.Users {
		users = append(users, userModel.ToDomain())
	}

	return users, nil
}
