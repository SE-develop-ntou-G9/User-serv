package repositories

import (
	"golangAPI/entity"

	Model "golangAPI/infrastructure/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByProviderID(provider, providerUserID string) (*entity.User, error) {
	var m Model.UserModel

	err := r.db.Where("provider = ? AND provider_user_id = ?", provider, providerUserID).First(&m).Error
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:             uint(m.Id),
		Provider:       m.Provider,
		ProviderUserID: m.ProviderUserID,
		Email:          m.Email,
		Name:           m.Name,
	}, nil
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	m := Model.UserModel{
		Provider:       user.Provider,
		ProviderUserID: user.ProviderUserID,
		Email:          user.Email,
		Name:           user.Name,
	}

	err := r.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	user.ID = uint(m.Id)

	return user, nil
}
