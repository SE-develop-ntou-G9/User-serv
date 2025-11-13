package entity

type User struct {
	ID             uint
	Provider       string
	ProviderUserID string
	Email          string
	Name           string
}

type UserRepository interface {
	FindAll() ([]User, error)
	Post(user User) error
}
