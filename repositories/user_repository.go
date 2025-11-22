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
		ID:             m.ID,
		Provider:       m.Provider,
		ProviderUserID: m.ProviderUserID,
		Email:          m.Email,
		Name:           m.Name,
	}, nil
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	m := Model.UserModel{
		ID:             user.ID,
		Provider:       user.Provider,
		ProviderUserID: user.ProviderUserID,
		Email:          user.Email,
		Name:           user.Name,
	}

	err := r.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	user.ID = m.ID

	return user, nil
}

func (r *userRepository) CreateDriver(driver *entity.Driver) (*entity.Driver, error) {
	m := Model.DriverModel{
		UserID:      driver.UserID,
		Name:        driver.Name,
		ContactInfo: driver.ContactInfo,
		ScooterType: driver.ScooterType,
		PlateNum:    driver.PlateNum,
	}

	err := r.db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (r *userRepository) EditUser(user *entity.User) (*entity.User, error) {
	m := Model.UserModel{
		ID:             user.ID,
		Provider:       user.Provider,
		ProviderUserID: user.ProviderUserID,
		Email:          user.Email,
		Name:           user.Name,
		PhoneNumber:    user.PhoneNumber,
	}

	err := r.db.Model(&Model.UserModel{}).Where("user_id = ?", user.ID).Updates(m).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(id string) (*entity.User, error) {
	var m Model.UserModel

	err := r.db.Where("user_id = ?", id).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:             m.ID,
		Provider:       m.Provider,
		ProviderUserID: m.ProviderUserID,
		Email:          m.Email,
		Name:           m.Name,
		PhoneNumber:    m.PhoneNumber,
	}, nil
}

func (r *userRepository) GetDriverByUserID(userID string) (*entity.Driver, error) {
	var m Model.DriverModel

	err := r.db.Where("user_id = ?", userID).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &entity.Driver{
		UserID:      m.UserID,
		Name:        m.Name,
		ContactInfo: m.ContactInfo,
		ScooterType: m.ScooterType,
		PlateNum:    m.PlateNum,
	}, nil
}
