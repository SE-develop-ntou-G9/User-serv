package usecase

import (
	"golangAPI/entity"
)

type UserRepository interface {
	FindByProviderID(provider, providerUserID string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	CreateDriver(driver *entity.Driver) (*entity.Driver, error)
}
