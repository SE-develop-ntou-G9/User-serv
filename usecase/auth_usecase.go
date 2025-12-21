package usecase

import (
	"errors"
	"golangAPI/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase struct {
	userRepo UserRepository
}

func NewAuthUsecase(r UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: r}
}

func (uc *AuthUsecase) LoginWithOAuth(provider string, gUser entity.User) (*entity.User, string, error) {
	u, err := uc.userRepo.FindByProviderID(provider, gUser.ProviderUserID)

	// 找不到才建立
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

	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 32 {
		// 32 bytes = 256 bits
		return nil, "", errors.New("JWT_SECRET must be at least 32 bytes")
	}

	role := "user"
	if u.Admin {
		role = "admin"
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"sub":   u.ID,
		"email": u.Email,
		"name":  u.Name,
		"role":  role,
		"iss":   "ntouber-user-serv",
		"iat":   now.Unix(),
		"exp":   now.Add(24 * time.Hour).Unix(),
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString([]byte(secret))
	if err != nil {
		return nil, "", err
	}

	return u, token, nil
}
