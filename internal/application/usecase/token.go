package usecase

import (
	"github.com/wang900115/LCB/internal/domain/entities"
	"github.com/wang900115/LCB/internal/domain/irepository"
)

type TokenUsecase struct {
	TokenRepo irepository.TokenRepository
}

func NewTokenUsecase(tokenRepo irepository.TokenRepository) *TokenUsecase {
	return &TokenUsecase{TokenRepo: tokenRepo}
}

func (t *TokenUsecase) CreateToken(tokenClaims entities.TokenClaims) (string, error) {
	return t.TokenRepo.CreateToken(tokenClaims)
}

func (t *TokenUsecase) ValidateToken(token string) (entities.TokenClaims, error) {
	return t.TokenRepo.ValidateToken(token)
}

func (t *TokenUsecase) DeleteToken(username, channelname string) error {
	return t.TokenRepo.DeleteToken(username, channelname)
}
