package irepository

import "github.com/wang900115/LCB/internal/domain/entities"

type TokenRepository interface {
	CreateToken(entities.TokenClaims) (string, error)
	ValidateToken(string) (entities.TokenClaims, error)
	DeleteToken(string, string) error
}
