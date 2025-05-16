package usecase

import (
	"github.com/wang900115/LCB/internal/domain/entities"
	"github.com/wang900115/LCB/internal/domain/irepository"
)

type UserUsecase struct {
	UserRepo irepository.UserRepository
}

func NewUserUsecase(userRepo irepository.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (u *UserUsecase) CreateUser(user entities.User) (entities.User, error) {
	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecase) DeleteUser(username string) (entities.User, error) {
	return u.UserRepo.DeleteUser(username)
}
