package model

type UserModel struct {
	ID             string `gorm:"column:user_id;primaryKey;autoIncrement:false"`
	Name           string `gorm:"column:user_name"`
	Provider       string `gorm:"column:provider"`
	ProviderUserID string `gorm:"column:provider_user_id"`
	Email          string `gorm:"column:user_email"`
	PhoneNumber    string `gorm:"column:phone_number"`
	AvatarURL      string `gorm:"column:avatar_url"`
	Admin          bool   `gorm:"column:is_admin;default:false" json:"is_admin"`

	Driver *DriverModel `gorm:"foreignKey:UserID;references:ID"`
}

func (UserModel) TableName() string {
	return "user_models"
}
