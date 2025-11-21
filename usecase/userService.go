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
