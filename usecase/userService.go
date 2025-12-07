package usecase

import (
	"golangAPI/entity"
)

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (uc *UserUsecase) CreateUser(u entity.User) (*entity.User, error) {
	return uc.repo.Create(&u)
}

func (uc *UserUsecase) CreateDriver(d entity.Driver) (*entity.Driver, error) {
	return uc.repo.CreateDriver(&d)
}

func (uc *UserUsecase) EditUser(u entity.User) (*entity.User, error) {
	return uc.repo.EditUser(&u)
}

func (uc *UserUsecase) GetUserByID(id string) (*entity.User, error) {
	return uc.repo.GetUserByID(id)
}

func (uc *UserUsecase) GetDriverByUserID(userID string) (*entity.Driver, error) {
	return uc.repo.GetDriverByUserID(userID)
}

func (uc *UserUsecase) GetAllUser() ([]entity.User, error) {
	return uc.repo.GetAllUser()
}

func (uc *UserUsecase) GetAllDriver() ([]entity.Driver, error) {
	return uc.repo.GetAllDriver()
}

func (uc *UserUsecase) DeleteAllUser() error {
	return uc.repo.DeleteAllUser()
}

func (uc *UserUsecase) DeleteUserByID(id string) error {
	return uc.repo.DeleteUserByID(id)
}

func (uc *UserUsecase) DeleteDriverByUserID(userID string) error {
	return uc.repo.DeleteDriverByUserID(userID)
}

func (uc *UserUsecase) EditDriver(d entity.Driver) (*entity.Driver, error) {
	return uc.repo.EditDriver(&d)
}
