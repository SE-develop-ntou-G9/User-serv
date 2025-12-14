package usecase

import (
	"golangAPI/entity"
	"mime/multipart"
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

type ImageRepository interface {
	UploadAvatarToS3(file multipart.File, fileName string, contentType string) (string, error)
	UploadLicenseToS3(file multipart.File, fileName string, contentType string) (string, error)
}

type NotifyRepository interface {
	Create(notify *entity.Notify) (*entity.Notify, error)
	GetByRecieverID(recieverID string) ([]entity.Notify, error)
	DeleteByRecieverID(id string) error
	DeleteByID(id string) error
}
