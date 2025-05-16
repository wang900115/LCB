package irepository

import "github.com/wang900115/LCB/internal/domain/entities"

type UserRepository interface {
	CreateUser(entities.User) (entities.User, error)
	DeleteUser(string) (entities.User, error)
}
