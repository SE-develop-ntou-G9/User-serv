package model

type UserModel struct {
	Id             int    `gorm:"column:user_id;primaryKey;"`
	Name           string `gorm:"column:user_name;"`
	Provider       string `gorm:"column:provider;"`
	ProviderUserID string `gorm:"column:provider_user_id;"`
	Email          string `gorm:"column:user_email;"`
}
