package usecase

import (
	"golangAPI/entity"
)

type UserRepository interface {
	FindByProviderID(provider, providerUserID string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	CreateDriver(driver *entity.Driver) (*entity.Driver, error)
	EditUser(user *entity.User) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	GetDriverByUserID(userID string) (*entity.Driver, error)
	DeleteAllUser() error
	DeleteUserByID(id string) error
	DeleteDriverByUserID(userID string) error
	EditDriver(driver *entity.Driver) (*entity.Driver, error)
}
