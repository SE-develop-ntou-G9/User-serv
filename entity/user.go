package entity

type User struct {
	ID             string
	Provider       string
	ProviderUserID string
	Email          string
	Name           string
}

type Driver struct {
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"driver_name" binding:"required"`
	ContactInfo string `json:"contact_info" binding:"required"`
	ScooterType string `json:"scooter_type" binding:"required"`
	PlateNum    string `json:"plate_num" binding:"required"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	Post(user User) error
}
