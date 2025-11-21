package usecase

import (
	"golangAPI/entity"
)

type AuthUsecase struct {
	userRepo UserRepository
}

func NewAuthUsecase(r UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: r}
}

func (uc *AuthUsecase) LoginWithOAuth(provider string, gUser entity.User) (*entity.User, string, error) {
	u, err := uc.userRepo.FindByProviderID(provider, gUser.ProviderUserID)
	if err != nil || u == nil {
		newUser := &entity.User{
			ID:             gUser.ID,
			Email:          gUser.Email,
			Name:           gUser.Name,
			Provider:       provider,
			ProviderUserID: gUser.ProviderUserID,
		}

		u, err = uc.userRepo.Create(newUser)
		if err != nil {
			return nil, "", err
		}
	}

	token := "mocked-jwt-token-for-user-" + u.ProviderUserID

	return u, token, nil

}
