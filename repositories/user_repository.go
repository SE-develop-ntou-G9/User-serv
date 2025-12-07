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
		Status:      driver.Status,
		DriverLicense: driver.DriverLicense,
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
		AvatarURL:      user.AvatarURL,
		Admin:          user.Admin,
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
		AvatarURL:      m.AvatarURL,
		Admin:          m.Admin,
	}, nil
}

func (r *userRepository) GetDriverByUserID(userID string) (*entity.Driver, error) {
	var m Model.DriverModel

	err := r.db.Where("user_id = ?", userID).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &entity.Driver{
		UserID:        m.UserID,
		Name:          m.Name,
		ContactInfo:   m.ContactInfo,
		ScooterType:   m.ScooterType,
		PlateNum:      m.PlateNum,
		DriverLicense: m.DriverLicense,
		Status:        m.Status,
	}, nil
}

func (r *userRepository) GetAllUser() ([]entity.User, error) {
	var models []Model.UserModel
	var users []entity.User

	err := r.db.Find(&models).Error
	if err != nil {
		return nil, err
	}

	for _, m := range models {
		users = append(users, entity.User{
			ID:             m.ID,
			Provider:       m.Provider,
			ProviderUserID: m.ProviderUserID,
			Email:          m.Email,
			Name:           m.Name,
			PhoneNumber:    m.PhoneNumber,
			AvatarURL:      m.AvatarURL,
		})
	}

	return users, nil
}

func (r *userRepository) GetAllDriver() ([]entity.Driver, error) {
	var models []Model.DriverModel
	var drivers []entity.Driver

	err := r.db.Find(&models).Error
	if err != nil {
		return nil, err
	}

	for _, m := range models {
		drivers = append(drivers, entity.Driver{
			UserID:        m.UserID,
			Name:          m.Name,
			ContactInfo:   m.ContactInfo,
			ScooterType:   m.ScooterType,
			PlateNum:      m.PlateNum,
			DriverLicense: m.DriverLicense,
      Status:        m.Status,
		})
	}

	return drivers, nil
}

func (r *userRepository) DeleteAllUser() error {
	err := r.db.Exec("DELETE FROM driver_models").Error
	if err != nil {
		return err
	}

	err = r.db.Exec("DELETE FROM user_models").Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUserByID(id string) error {
	err := r.db.Where("user_id = ?", id).Delete(&Model.DriverModel{}).Error
	if err != nil {
		return err
	}
	err = r.db.Where("user_id = ?", id).Delete(&Model.UserModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteDriverByUserID(userID string) error {
	err := r.db.Where("user_id = ?", userID).Delete(&Model.DriverModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) EditDriver(driver *entity.Driver) (*entity.Driver, error) {
	m := Model.DriverModel{
		UserID:        driver.UserID,
		Name:          driver.Name,
		ContactInfo:   driver.ContactInfo,
		ScooterType:   driver.ScooterType,
		PlateNum:      driver.PlateNum,
		DriverLicense: driver.DriverLicense,
		Status:        driver.Status,
	}
	err := r.db.Model(&Model.DriverModel{}).Where("user_id = ?", driver.UserID).Updates(m).Error
	if err != nil {
		return nil, err
	}

	return driver, nil
}
