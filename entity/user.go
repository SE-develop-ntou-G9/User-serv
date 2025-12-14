package entity

type User struct {
	ID             string
	Provider       string
	ProviderUserID string
	Email          string
	Name           string
	PhoneNumber    string
	AvatarURL      string
	Admin          bool
}

type Driver struct {
	UserID        string `json:"user_id" binding:"required"`
	Name          string `json:"driver_name" binding:"required"`
	ContactInfo   string `json:"contact_info" binding:"required"`
	ScooterType   string `json:"scooter_type" binding:"required"`
	PlateNum      string `json:"plate_num" binding:"required"`
	DriverLicense string `json:"driver_license" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

type Notify struct {
	ID         string `gorm:"primaryKey;type:char(36)"`
	RecieverID string
	SenderID   string
	Message    string
	Status     string
	TimeStamp  string `gorm:"column:time_stamp"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	Post(user User) error
}
