package usecase

import (
	"golangAPI/entity"

	"github.com/markbates/goth"
)

type UserRepository interface {
	FindByProviderID(provider, providerUserID string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}

type AuthUsecase struct {
	userRepo UserRepository
}

func NewAuthUsecase(r UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: r}
}

func (uc *AuthUsecase) LoginWithOAuth(provider string, gUser goth.User) (*entity.User, string, error) {
	u, err := uc.userRepo.FindByProviderID(provider, gUser.UserID)
	if err != nil || u == nil {
		newUser := &entity.User{
			Email:          gUser.Email,
			Name:           gUser.Name,
			Provider:       provider,
			ProviderUserID: gUser.UserID,
		}

		u, err = uc.userRepo.Create(newUser)
		if err != nil {
			return nil, "", err
		}
	}

	token := "mocked-jwt-token-for-user-" + u.ProviderUserID

	return u, token, nil

}
