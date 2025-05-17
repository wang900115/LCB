package repository

import (
	"github.com/wang900115/LCB/internal/adapter/gorm/model"
	"github.com/wang900115/LCB/internal/domain/entities"
	"github.com/wang900115/LCB/internal/domain/irepository"
	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) irepository.UserRepository {
	return &UserRepository{gorm: gorm}
}

func (r *UserRepository) CreateUser(user entities.User) (entities.User, error) {
	userModel := model.User{
		Name: user.Name,
		Channel: model.Channel{
			Model: gorm.Model{ID: user.ChannelID},
		},
	}

	if err := r.gorm.Create(&userModel).Error; err != nil {
		return entities.User{}, err
	}

	return userModel.ToDomain(), nil
}

func (r *UserRepository) DeleteUser(username string) (entities.User, error) {
	var user model.User
	if err := r.gorm.Where("name = ?", username).First(&user).Error; err != nil {
		return entities.User{}, err
	}

	if err := r.gorm.Unscoped().Delete(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user.ToDomain(), nil
}
