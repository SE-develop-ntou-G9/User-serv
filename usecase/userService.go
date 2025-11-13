package usecase

import (
	"golangAPI/entity"
)

type UserUsecase struct {
	repo UserRepository
}

func (uc *UserUsecase) CreateUser(u entity.User) (*entity.User, error) {
	return uc.repo.Create(&u)
}
