package model

type DriverModel struct {
	UserID        string `gorm:"column:user_id;primaryKey;autoIncrement:false"`
	Name          string `gorm:"column:driver_name"`
	ContactInfo   string `gorm:"column:contact_info"`
	ScooterType   string `gorm:"column:scooter_type"`
	PlateNum      string `gorm:"column:plate_num"`
	DriverLicense string `gorm:"column:driver_license"`
	Status        string `gorm:"column:status"`

	User UserModel `gorm:"foreignKey:UserID;references:ID"`
}

func (DriverModel) TableName() string {
	return "driver_models"
}
